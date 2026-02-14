package kvstore

import (
	"context"
	"sync"
	"time"
)

// Store defines the interface for key-value storage operations
type Store interface {
	// Get retrieves a value by key. Returns (value, found)
	Get(ctx context.Context, key string) (interface{}, bool)

	// Set stores a value with a TTL duration
	Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error

	// Increment atomically increments a counter and returns the new value
	Increment(ctx context.Context, key string, value int64, ttl time.Duration) (int64, error)

	// Delete removes a key
	Delete(ctx context.Context, key string) error

	// Clear removes all keys
	Clear(ctx context.Context) error

	// Close gracefully shuts down the store
	Close() error
	// Keys returns a slice of keys present in the store (useful for small-scale scanning)
	Keys(ctx context.Context) []string
}

// entry represents a stored value with expiration time
type entry struct {
	value     interface{}
	expiresAt time.Time
}

// MemoryStore is a thread-safe in-memory implementation of Store
type MemoryStore struct {
	mu                sync.RWMutex
	data              map[string]entry
	cleanupInterval   time.Duration
	cleanupTicker     *time.Ticker
	stopCleanup       chan struct{}
	cleanupEnabled    bool
}

// NewMemoryStore creates a new in-memory key-value store with configurable cleanup
// cleanupInterval: how often to clean up expired entries (0 = no periodic cleanup, only lazy)
func NewMemoryStore(cleanupInterval time.Duration) *MemoryStore {
	ms := &MemoryStore{
		data:            make(map[string]entry),
		cleanupInterval: cleanupInterval,
		stopCleanup:     make(chan struct{}),
		cleanupEnabled:  cleanupInterval > 0,
	}

	// Start periodic cleanup goroutine if enabled
	if ms.cleanupEnabled {
		ms.startCleanupRoutine()
	}

	return ms
}

// Get retrieves a value by key with lazy expiration check
func (ms *MemoryStore) Get(ctx context.Context, key string) (interface{}, bool) {
	ms.mu.RLock()
	entry, exists := ms.data[key]
	ms.mu.RUnlock()

	if !exists {
		return nil, false
	}

	// Lazy expiration check
	if time.Now().After(entry.expiresAt) {
		ms.mu.Lock()
		delete(ms.data, key)
		ms.mu.Unlock()
		return nil, false
	}

	return entry.value, true
}

// Set stores a value with TTL
func (ms *MemoryStore) Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	ms.mu.Lock()
	defer ms.mu.Unlock()

	ms.data[key] = entry{
		value:     value,
		expiresAt: time.Now().Add(ttl),
	}

	return nil
}

// Increment atomically increments a counter
func (ms *MemoryStore) Increment(ctx context.Context, key string, value int64, ttl time.Duration) (int64, error) {
	ms.mu.Lock()
	defer ms.mu.Unlock()

	existingEntry, exists := ms.data[key]

	// Lazy expiration check
	if exists && time.Now().After(existingEntry.expiresAt) {
		exists = false
	}

	var currentValue int64
	if exists {
		if counter, ok := existingEntry.value.(int64); ok {
			currentValue = counter
		}
	}

	newValue := currentValue + value
	ms.data[key] = entry{
		value:     newValue,
		expiresAt: time.Now().Add(ttl),
	}

	return newValue, nil
}

// Delete removes a key
func (ms *MemoryStore) Delete(ctx context.Context, key string) error {
	ms.mu.Lock()
	defer ms.mu.Unlock()

	delete(ms.data, key)
	return nil
}

// Clear removes all keys
func (ms *MemoryStore) Clear(ctx context.Context) error {
	ms.mu.Lock()
	defer ms.mu.Unlock()

	ms.data = make(map[string]entry)
	return nil
}

// Keys returns a list of all keys currently stored (not recommended for large stores)
func (ms *MemoryStore) Keys(ctx context.Context) []string {
	ms.mu.RLock()
	defer ms.mu.RUnlock()
	keys := make([]string, 0, len(ms.data))
	for k := range ms.data {
		keys = append(keys, k)
	}
	return keys
}

// startCleanupRoutine starts the periodic cleanup goroutine
func (ms *MemoryStore) startCleanupRoutine() {
	ms.cleanupTicker = time.NewTicker(ms.cleanupInterval)

	go func() {
		for {
			select {
			case <-ms.cleanupTicker.C:
				ms.cleanupExpired()
			case <-ms.stopCleanup:
				ms.cleanupTicker.Stop()
				return
			}
		}
	}()
}

// cleanupExpired removes all expired entries
func (ms *MemoryStore) cleanupExpired() {
	ms.mu.Lock()
	defer ms.mu.Unlock()

	now := time.Now()
	for key, entry := range ms.data {
		if now.After(entry.expiresAt) {
			delete(ms.data, key)
		}
	}
}

// Close gracefully shuts down the store
func (ms *MemoryStore) Close() error {
	if ms.cleanupEnabled && ms.cleanupTicker != nil {
		close(ms.stopCleanup)
		// Give cleanup goroutine time to exit
		time.Sleep(10 * time.Millisecond)
	}
	return nil
}

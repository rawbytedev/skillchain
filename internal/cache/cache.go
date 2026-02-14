package cache

import (
	"context"
	"time"

	"moltket/internal/kvstore"
)

// Client wraps the KVStore interface for caching operations
type Client struct {
	store kvstore.Store
}

// NewClient creates a new cache client with the provided store
func NewClient(store kvstore.Store) *Client {
	return &Client{
		store: store,
	}
}

// NewKVStore creates a cache client backed by an in-memory KV store (convenience helper for tests)
func NewKVStore() *Client {
	ms := kvstore.NewMemoryStore(0)
	return NewClient(ms)
}

// Get retrieves a cached value
func (c *Client) Get(ctx context.Context, key string) (interface{}, bool) {
	return c.store.Get(ctx, key)
}

// Set stores a value with TTL
func (c *Client) Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	return c.store.Set(ctx, key, value, ttl)
}

// Increment atomically increments a counter
func (c *Client) Increment(ctx context.Context, key string, value int64, ttl time.Duration) (int64, error) {
	return c.store.Increment(ctx, key, value, ttl)
}

// Delete removes a cached value
func (c *Client) Delete(ctx context.Context, key string) error {
	return c.store.Delete(ctx, key)
}

// GetStore returns the underlying KVStore (for advanced usage)
func (c *Client) GetStore() kvstore.Store {
	return c.store
}

// Close gracefully shuts down the cache
func (c *Client) Close() error {
	return c.store.Close()
}

// Clear removes all keys from the underlying store
func (c *Client) Clear(ctx context.Context) error {
	return c.store.Clear(ctx)
}

// Keys returns all keys from the underlying store (useful for small-scale scanning)
func (c *Client) Keys(ctx context.Context) []string {
	return c.store.Keys(ctx)
}

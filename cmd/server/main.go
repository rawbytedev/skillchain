package main

import (
	"context"
	"log"
	"math/big"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"moltket/config"
	"moltket/internal/api"
	"moltket/internal/auth"
	"moltket/internal/blockchain"
	"moltket/internal/cache"
	"moltket/internal/core"

	"github.com/ethereum/go-ethereum/common"
)

func main() {
	// Load configuration
	cfg := config.Load()

	log.Printf("Starting SkillChain Verification Service in %s mode", cfg.Env)

	// Initialize dependencies
	kvStore := cache.NewKVStore()
	defer kvStore.Close()

	// Initialize blockchain client if enabled
	var bcClient *blockchain.Client
	if cfg.EnableBlockchain {
		var err error
		bcClient, err = blockchain.NewClient(cfg.EthNodeURL)
		if err != nil {
			log.Printf("Warning: Failed to connect to blockchain: %v", err)
			log.Println("Continuing in offline mode...")
		} else {
			defer bcClient.Close()
		}
	}

	// Initialize EIP-712 signer for licenses
	chainID := big.NewInt(cfg.ChainID)
	verifyingContract := common.HexToAddress(cfg.LicenseNFTAddress)

	signer, err := auth.NewSigner(cfg.SignerPrivateKey, chainID, verifyingContract)
	if err != nil {
		log.Fatalf("Failed to create signer: %v", err)
	}

	// Initialize vote service for batch processor
	voteService := core.NewVoteService(cfg, kvStore, signer)
	batchProcessor := core.NewBatchProcessor(voteService, kvStore, 5*time.Minute)

	// Create and start server with all components
	server := api.NewServer(cfg, kvStore, bcClient, voteService)

	// Start batch processor (runs automatically every 5 minutes)
	ctx := context.Background()
	batchProcessor.Start(ctx)
	defer batchProcessor.Stop()
	// Start server in a goroutine
	go func() {
		if err := server.Start(":" + cfg.ServerPort); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed: %v", err)
		}
	}()

	log.Printf("Server started on port %s", cfg.ServerPort)
	log.Printf("Batch processor running every 5 minutes")

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited properly")
}

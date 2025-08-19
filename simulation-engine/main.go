package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"simengine/api"
	"syscall"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading env file")
	}
	
    addr := os.Getenv("PORT")
    server := &http.Server{
        Addr:         addr,
        Handler:      api.NewRouter(),
        ReadTimeout:  15 * time.Second,
        WriteTimeout: 15 * time.Second,
        IdleTimeout:  60 * time.Second,
    }

    stop := make(chan os.Signal, 1)
    signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

    go func() {
        log.Printf("Simulator Engine is running at %s\n", addr)
        if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            log.Fatalf("Failed to start server: %v", err)
        }
    }()

    <-stop

    shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    if err := server.Shutdown(shutdownCtx); err != nil {
        log.Fatalf("Shutdown failed: %v", err)
    }

    log.Println("Engine stopped cleanly")
}
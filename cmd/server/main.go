package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/zadescoxp/GoCord.git/internal/config"
	"github.com/zadescoxp/GoCord.git/internal/database"
	"github.com/zadescoxp/GoCord.git/internal/routes"
)

func main() {
	cfg := config.LoadConfig()
	db := database.Connect(cfg.DBURL)
	router := routes.SetupRouter()

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.Port),
		Handler: router,
	}

	log.Printf("🚀 Server running on port %s\n", cfg.Port)
	log.Println(db)

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}

	log.Println("Server exited properly.")
}

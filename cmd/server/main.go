package main

import (
	"fmt"
	"log"

	"github.com/zadescoxp/GoCord/internal/config"
	"github.com/zadescoxp/GoCord/internal/routes"
)

func main() {
	cfg := config.LoadConfig()
	router := routes.SetupRouter()

	log.Printf("🚀 Server running on port %s\n", cfg.Port)

	err := router.Run(fmt.Sprintf(":%s", cfg.Port))

	if err != nil {
		log.Fatal(err)
	}
}

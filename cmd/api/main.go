package main

import (
	"fmt"
	"log"

	"github.com/Sina-Mahmoodmoradi/ecommerce-backend/internal/config"
)

func main() {
	cfg, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	fmt.Printf("Running %s on port %s in %s mode\n", cfg.AppName, cfg.AppPort, cfg.AppEnv)
}

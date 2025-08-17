package main

import (
	"fmt"
	"log"

	"github.com/Sina-Mahmoodmoradi/ecommerce-backend/internal/config"
	"github.com/Sina-Mahmoodmoradi/ecommerce-backend/internal/http/middleware"
	"github.com/Sina-Mahmoodmoradi/ecommerce-backend/internal/pkg/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	cfg, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	logger := logger.New(cfg.AppEnv)
	defer logger.Sync()

	route := gin.New()
	route.Use(middleware.RequestID())
	route.Use(middleware.GinLogger(logger))
	route.Use(gin.Recovery())

	route.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})
	logger.Info(fmt.Sprintf("Starting server on port %s", cfg.AppPort))
	if err := route.Run(":" + cfg.AppPort); err != nil {
		logger.Fatal("Failed to start server", zap.Error(err))
	}
}

package config

import (
	"fmt"
	// "log"
	"strings"

	"github.com/spf13/viper"
)

// Config holds all application configuration.
type Config struct {
	AppName string `mapstructure:"APP_NAME"`
	AppEnv  string `mapstructure:"APP_ENV"`
	AppPort string `mapstructure:"APP_PORT"`

	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     int    `mapstructure:"DB_PORT"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBName     string `mapstructure:"DB_NAME"`
	DBSSLMode  string `mapstructure:"DB_SSLMODE"`

	RedisHost string `mapstructure:"REDIS_HOST"`
	RedisPort int    `mapstructure:"REDIS_PORT"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (Config, error) {
	var config Config

	viper.SetConfigName("config") // config.yaml or config.json
	viper.AddConfigPath(path)
	viper.SetConfigType("yaml")
	viper.AutomaticEnv() // Read matching env vars
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// Default values
	viper.SetDefault("APP_NAME", "ecommerce-backend")
	viper.SetDefault("APP_ENV", "development")
	viper.SetDefault("APP_PORT", "8080")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("No config file found, relying on environment variables...")
	}

	if err := viper.Unmarshal(&config); err != nil {
		return config, fmt.Errorf("unable to decode config into struct: %w", err)
	}

	return config, nil
}

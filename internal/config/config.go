package config

import (
	"os"
)

// Config holds application configuration
type Config struct {
	// Add configuration fields as needed
	JWTSecret string
	DBHost    string
	DBPort    string
	DBUser    string
	DBPass    string
	DBName    string
}

// Load reads configuration from environment variables
func Load() (*Config, error) {
	return &Config{
		JWTSecret: getEnvOrDefault("JWT_SECRET", "your-secret-key"),
		DBHost:    getEnvOrDefault("DB_HOST", "localhost"),
		DBPort:    getEnvOrDefault("DB_PORT", "5432"),
		DBUser:    getEnvOrDefault("DB_USER", "postgres"),
		DBPass:    getEnvOrDefault("DB_PASS", "postgres"),
		DBName:    getEnvOrDefault("DB_NAME", "mealsync"),
	}, nil
}

// getEnvOrDefault returns environment variable value or default if not set
func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

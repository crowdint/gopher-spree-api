package configs

import (
	"os"
)

const (
	DB_DEBUG      = "DATABASE_DEBUG"
	DB_ENGINE     = "DATABASE_ENGINE"
	DB_IDLE_CONNS = "MAX_IDLE_CONNS"
	DB_OPEN_CONNS = "MAX_OPEN_CONNS"
	DB_URL        = "DATABASE_URL"

	SPREE_URL = "SPREE_URL"
	SPREE_NS  = "SPREE_NAMESPACE"
)

var defaultSettings = map[string]string{
	// Database Settings
	DB_DEBUG:      getEnvOrDefault(DB_DEBUG, "true"),
	DB_ENGINE:     getEnvOrDefault(DB_ENGINE, "postgres"),
	DB_URL:        getEnvOrDefault(DB_URL, "dbname=spree_dev sslmode=disable"),
	DB_IDLE_CONNS: getEnvOrDefault(DB_IDLE_CONNS, "10"),
	DB_OPEN_CONNS: getEnvOrDefault(DB_OPEN_CONNS, "100"),

	// Spree Settings
	SPREE_URL: getEnvOrDefault(SPREE_URL, "http://localhost:5100"),
	SPREE_NS:  getEnvOrDefault(SPREE_NS, "store"),
}

func Get(key string) string {
	return defaultSettings[key]
}

func getEnvOrDefault(key, defaultValue string) string {
	if envValue := os.Getenv(key); envValue != "" {
		return envValue
	}
	return defaultValue
}

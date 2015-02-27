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

	MEMCACHED_URL = "MEMCACHED_URL"

	SPREE_URL = "SPREE_URL"
	SPREE_NS  = "SPREE_NAMESPACE"

	SPREE_ASSET_PATH     = "SPREE_ASSET_PATH"
	SPREE_ASSET_HOST     = "SPREE_ASSET_HOST"
	SPREE_DEFAULT_STYLES = "SPREE_DEFAULT_STYLES"

	PER_PAGE = "PER_PAGE"
)

var defaultSettings = map[string]string{
	// Database Settings
	DB_DEBUG:      getEnvOrDefault(DB_DEBUG, "false"),
	DB_ENGINE:     getEnvOrDefault(DB_ENGINE, "postgres"),
	DB_URL:        getEnvOrDefault(DB_URL, "dbname=spree_dev sslmode=disable"),
	DB_IDLE_CONNS: getEnvOrDefault(DB_IDLE_CONNS, "2"),
	DB_OPEN_CONNS: getEnvOrDefault(DB_OPEN_CONNS, "5"),

	// Memcached Settings
	MEMCACHED_URL: getEnvOrDefault(MEMCACHED_URL, "127.0.0.1:11211"),

	// Spree Settings
	SPREE_URL: getEnvOrDefault(SPREE_URL, "http://localhost:5100"),
	SPREE_NS:  getEnvOrDefault(SPREE_NS, ""),
	//Spree Assets Default Settings
	SPREE_ASSET_PATH:     getEnvOrDefault(SPREE_ASSET_PATH, ":host/spree/products/:id/:style/:filename"),
	SPREE_ASSET_HOST:     getEnvOrDefault(SPREE_ASSET_HOST, ""),
	SPREE_DEFAULT_STYLES: getEnvOrDefault(SPREE_DEFAULT_STYLES, "mini,small,product,large"),

	//Response Settings
	PER_PAGE: getEnvOrDefault(PER_PAGE, "10"),
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

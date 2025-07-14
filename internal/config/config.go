package config

import (
	"os"
	"strconv"
)

type Config struct {
	PGUSER string
	PGPASS string
	PGHOST string
	PGPORT int
	PGDB   string
}

func New() *Config {
	return &Config{
		PGUSER: getEnv("PGUSER", ""),
		PGPASS: getEnv("PGPASS", ""),
		PGHOST: getEnv("PGHOST", "localhost"),
		PGPORT: getEnvAsInt("PGPORT", 5432),
		PGDB:   getEnv("PGDB", "swim_help"),
	}
}

func getEnv(key string, defaultValue string) string {
	value, exists := os.LookupEnv(key)

	if exists {
		return value
	}

	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	valueStr := getEnv(key, "")

	if value, errConv := strconv.Atoi(valueStr); errConv == nil {
		return value
	}

	return defaultValue
}

package utils

import (
	"fmt"
	"lido-core/v1/pkg/configs"
	"os"
)

// ConnectionURLBuilder func for building URL connection.
func ConnectionURLBuilder(n string) (string, error) {
	// Define URL to connection.
	var url string

	// Switch given names.
	switch n {
	case "mongo":
		// URL for Mysql connection.
		url = fmt.Sprintf(
			"mongodb://%s:%s@%s:%s/%s?authSource=admin",
			configs.DbUser,
			configs.DbPassword,
			configs.DbHost,
			configs.DbPort,
			configs.DbName,
		)
	case "redis":
		// URL for Redis connection.
		url = fmt.Sprintf(
			"%s:%s",
			os.Getenv("REDIS_HOST"),
			os.Getenv("REDIS_PORT"),
		)
	case "fiber":
		// URL for Fiber connection.
		url = fmt.Sprintf(
			"%s:%s",
			os.Getenv("SERVER_HOST"),
			os.Getenv("SERVER_PORT"),
		)
	default:
		// Return error message.
		return "", fmt.Errorf("connection name '%v' is not supported", n)
	}

	// Return connection URL.
	return url, nil
}

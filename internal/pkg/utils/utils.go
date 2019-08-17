package utils

import "os"

// GetEnvOrDefault returns default value if environment variable is not set
func GetEnvOrDefault(key string, def string) (val string) {
	val = os.Getenv(key)

	if val == "" {
		val = def
	}

	return
}

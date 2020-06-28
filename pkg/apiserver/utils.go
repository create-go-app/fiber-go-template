package apiserver

import "os"

// GetEnv ...
func GetEnv(key, fallback string) string {
	// Check, if the environment variable exists
	if value, ok := os.LookupEnv(key); ok {
		// Return value
		return value
	}

	// Return fallback value
	return fallback
}

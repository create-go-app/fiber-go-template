package apiserver

import "os"

// GetEnv function for getting environment variables values,
// or return fallback
func GetEnv(key, fallback string) string {
	// Check, if the environment variable exists
	if value, ok := os.LookupEnv(key); ok {
		// Return value
		return value
	}

	// Return fallback value
	return fallback
}

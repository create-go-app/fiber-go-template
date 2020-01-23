package apiserver

import (
	"os"

	"go.uber.org/zap"
)

// IsError method for check error and show message
func (s *APIServer) IsError(err error) {
	// If got error
	if err != nil {
		// Show error in logs
		s.logger.Error("Error", zap.Error(err))

		// Exit with status 1
		os.Exit(1)
	}
}

package apiserver

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger function for init new zap logger instance
func Logger(config *Config) *zap.Logger {
	// Define log level
	level := zap.NewAtomicLevel()

	// Set log level from .env file
	switch config.Logging.Level {
	case "debug":
		level.SetLevel(zap.DebugLevel)
	case "warn":
		level.SetLevel(zap.WarnLevel)
	case "error":
		level.SetLevel(zap.ErrorLevel)
	case "fatal":
		level.SetLevel(zap.FatalLevel)
	case "panic":
		level.SetLevel(zap.PanicLevel)
	default:
		level.SetLevel(zap.InfoLevel)
	}

	// Create new zap logger config
	encoderCfg := zap.NewProductionEncoderConfig()

	// Formated timestamp in the output.
	encoderCfg.EncodeTime = zapcore.RFC3339TimeEncoder

	// Create new zap logger
	logger := zap.New(zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderCfg),
		zapcore.Lock(os.Stdout),
		level,
	))
	defer logger.Sync()

	return logger
}

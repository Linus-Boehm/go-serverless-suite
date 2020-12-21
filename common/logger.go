package common

import (
	"context"
	"os"

	"github.com/rs/zerolog"
)

const envLogMode = "LOG_MODE"
const envLogModeProduction = "production"

var instance zerolog.Logger
var initiated = false

func GetDefaultLogger() *zerolog.Logger {
	if !initiated {
		instance = zerolog.New(os.Stderr).With().Timestamp().Logger()

		if os.Getenv(envLogMode) == envLogModeProduction {
			// InfoLevel and above to standard error as JSON
			zerolog.SetGlobalLevel(zerolog.InfoLevel)
			instance.Info().Msg("Initiating logger with InfoLevel")
		} else {
			zerolog.SetGlobalLevel(zerolog.DebugLevel)
			instance.Info().Msg("Initiating logger with DebugLevel")
		}
		initiated = true
	}
	return &instance
}

func ConfigureLogger(cfg Configer) error {
	if !cfg.IsDebug() && os.Getenv(envLogMode) == "" {
		return os.Setenv(envLogMode, envLogModeProduction)
	}
	return nil
}

func WithContext(ctx *context.Context) *zerolog.Logger {
	logger := GetDefaultLogger()
	//TODO read path, request id etc
	return logger
}

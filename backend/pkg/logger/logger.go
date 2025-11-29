package logger

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Init initializes the global logger based on configuration
func Init(level string, format string) {
	// Set time format
	zerolog.TimeFieldFormat = time.RFC3339

	// Set log level
	logLevel, err := zerolog.ParseLevel(level)
	if err != nil {
		logLevel = zerolog.InfoLevel
	}
	zerolog.SetGlobalLevel(logLevel)

	// Set output format
	if format == "console" {
		// Pretty console output for development
		log.Logger = zerolog.New(zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: time.RFC3339,
		}).With().Timestamp().Logger()
	} else {
		// JSON output for production
		log.Logger = zerolog.New(os.Stdout).With().Timestamp().Logger()
	}

	log.Info().
		Str("level", logLevel.String()).
		Str("format", format).
		Msg("Logger initialized")
}


// Info logs an info message
func Info() *zerolog.Event {
	return log.Info()
}

// Error logs an error message
func Error() *zerolog.Event {
	return log.Error()
}

// Warn logs a warning message
func Warn() *zerolog.Event {
	return log.Warn()
}

// Fatal logs a fatal message
func Fatal() *zerolog.Event {
	return log.Fatal()
}

// Debug logs a debug message
func Debug() *zerolog.Event {
	return log.Debug()
}


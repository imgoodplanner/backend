package logger

import (
	"context"
	"os"
	"path/filepath"
	"time"

	"github.com/rs/zerolog"
)

type CtxKey struct{}

// SetLogger returns ctx with Logger set inside it using CtxKey
func SetLogger(ctx context.Context, logger *zerolog.Logger) context.Context {
	return context.WithValue(ctx, CtxKey{}, logger)
}

func FromCtx(ctx context.Context) *zerolog.Logger {
	return ctx.Value(CtxKey{}).(*zerolog.Logger)
}

func ZerologLogger(level string) zerolog.Logger {
	serviceName := filepath.Base(os.Args[0])
	serviceVersion := os.Getenv("VERSION")
	zerolog.TimeFieldFormat = time.RFC3339
	l := zerolog.
		New(os.Stdout).
		With().
		Timestamp().
		Str("service_name", serviceName).
		Str("service_version", serviceVersion).
		Caller().
		Logger()

	zerolog.SetGlobalLevel(GetLogLevel(level))

	// do not log in json in development environment
	if os.Getenv("ENV") == "development" {
		l = l.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	return l
}

func GetLogLevel(level string) zerolog.Level {
	switch level {
	case "debug":
		return zerolog.DebugLevel
	case "info":
		return zerolog.InfoLevel
	case "warning":
		return zerolog.WarnLevel
	case "error":
		return zerolog.ErrorLevel
	case "fatal":
		return zerolog.FatalLevel
	case "panic":
		return zerolog.PanicLevel
	default:
		// defaulting to info
		return zerolog.InfoLevel
	}

}

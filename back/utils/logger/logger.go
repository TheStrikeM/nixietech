package logger

import (
	"log/slog"
	"nixietech/utils/logger/handler"
	"os"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

type Logger struct{}

func SetupLogger(env string, withParams bool) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = setupPrettySlog()
	case envDev:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	if withParams {
		log = log.With(slog.String("env", env))
	}

	slog.SetDefault(log)
	return log
}

func setupPrettySlog() *slog.Logger {
	opts := handler.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	newHandler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(newHandler)
}

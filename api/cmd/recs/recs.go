package main

import (
	apppkg "krr/internal/app"
	"krr/internal/config"
	"krr/internal/logger"
	"log/slog"
)

func main() {
	cfg := config.MustLoad()
	log := logger.New(cfg.Env)
	app := apppkg.New(log, cfg.Env, cfg.Port)

	err := app.Server.Serve()
	log.Error("application has stopped: %s", slog.Any("error", err))
}

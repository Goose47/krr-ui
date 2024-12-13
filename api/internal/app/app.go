// Package app defines application model.
package app

import (
	"context"
	serverapp "krr/internal/app/server"
	"krr/internal/controllers"
	"krr/internal/server"
	"krr/internal/services"
	"log/slog"
)

type clearer interface {
	Clear(ctx context.Context)
}

// App represents application.
type App struct {
	Server *serverapp.Server
	Cache  clearer
}

// New creates all dependencies for App and returns new App instance.
func New(
	log *slog.Logger,
	env string,
	port int,
) *App {
	krrService := services.NewKRRService(log)
	recsService := services.NewRecsService(log, krrService)

	recsCon := controllers.NewRecsController(log, recsService)

	router := server.NewRouter(log, env, recsCon)
	serverApp := serverapp.New(log, port, router)

	return &App{
		Server: serverApp,
	}
}

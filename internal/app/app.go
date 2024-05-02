package app

import (
	"log/slog"
	appgrpc "profile/internal/app/app"
)

type App struct {
	GRPC *appgrpc.App
}

func New(log *slog.Logger, port int) *App {
	return &App{GRPC: nil}
}

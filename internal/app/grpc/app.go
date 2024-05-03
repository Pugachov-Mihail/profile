package appgrpc

import (
	"fmt"
	"google.golang.org/grpc"
	"log/slog"
	"net"
	"profile/internal/grpc/profile"
)

type App struct {
	log        *slog.Logger
	grpcServer *grpc.Server
	port       int
}

func New(log *slog.Logger, port int) *App {
	grpcServer := grpc.NewServer()

	profile.ServerApi(grpcServer)
	return &App{
		log:        log,
		grpcServer: grpcServer,
		port:       port,
	}
}

func (a *App) run() error {
	log := a.log.With("Auth service Running", a.port)

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))

	if err != nil {
		log.Error("Error Run Auth service")
		return fmt.Errorf("%s: %w", "Auth Run", err)
	}

	log.Info("Auth Run", slog.String("addr", l.Addr().String()))

	if err := a.grpcServer.Serve(l); err != nil {
		log.Error("Error Run GRPC Auth service")
		return fmt.Errorf("%s: %w", "Auth Run", err)
	}

	return nil
}

func (a *App) Stop() {
	a.log.With("Auth service Spoping", a.port).
		Info("Auth GRPC stop")

	a.grpcServer.GracefulStop()
}

func (a *App) MustRun() {
	if err := a.run(); err != nil {
		panic(err)
	}
}

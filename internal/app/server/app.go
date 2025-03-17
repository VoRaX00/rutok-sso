package server

import (
	"fmt"
	"google.golang.org/grpc"
	"log/slog"
	"net"
	"rutok-sso/internal/grpc/auth"
)

type App struct {
	log        *slog.Logger
	gRPCServer *grpc.Server
	port       int
}

func NewApp(log *slog.Logger, server auth.Auth, port int) *App {
	gRPCServer := grpc.NewServer()
	auth.Register(gRPCServer, server)

	return &App{
		log:        log,
		gRPCServer: grpc.NewServer(),
	}
}

func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(a.Run())
	}
}

func (a *App) Run() error {
	const op = "server.Run"
	log := a.log.With(
		slog.String("op", op),
	)

	log.Info("starting server")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if err = a.gRPCServer.Serve(lis); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	log.Info("server started")
	return nil
}

func (a *App) Stop() {
	const op = "server.Stop"
	a.log.With(
		slog.String("op", op),
	).Info("stopping server")

	a.gRPCServer.GracefulStop()
}

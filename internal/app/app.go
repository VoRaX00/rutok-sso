package app

import (
	"log/slog"
	grpcapp "rutok-sso/internal/app/server"
	"rutok-sso/internal/service/auth"
	"time"
)

type App struct {
	GRPCServer *grpcapp.App
}

func NewApp(log *slog.Logger, grpcPort int, storagePath string, tokenTTL time.Duration) *App {
	authService := auth.NewAuth(log, tokenTTL)
	grpcApp := grpcapp.NewApp(log, authService, grpcPort)
	return &App{
		GRPCServer: grpcApp,
	}
}

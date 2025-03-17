package auth

import (
	"log/slog"
	"time"
)

type Auth struct {
	log      *slog.Logger
	tokenTTL time.Duration
}

func NewAuth(log *slog.Logger, tokenTTL time.Duration) *Auth {
	return &Auth{
		log:      log,
		tokenTTL: tokenTTL,
	}
}

package auth

import (
	"context"
	"github.com/google/uuid"
	"log/slog"
	"time"
)

type UserSaver interface {
	Save(ctx context.Context, email string, password []byte) (uuid.UUID, error)
}

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

func (a *Auth) Login(ctx context.Context, email, password string, appId int32) (string, error) {
	panic("implement me")
}

func (a *Auth) Register(ctx context.Context, email, login, password string) (string, error) {
	panic("implement me")
}

func (a *Auth) IsAdmin(ctx context.Context, email string) (bool, error) {
	panic("implement me")
}

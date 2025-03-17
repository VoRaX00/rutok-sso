package auth

import (
	"context"
	ssov1 "github.com/VoRaX00/rutok-sso-proto/gen/go/sso"
	"google.golang.org/grpc"
)

type Auth interface {
	Login(ctx context.Context, email, password string, appId int32) (string, error)
	Register(ctx context.Context, email, login, password string) (string, error)
	IsAdmin(ctx context.Context, userId string) (bool, error)
}

type serverAPI struct {
	ssov1.UnimplementedAuthServer
	auth Auth
}

func Register(gRPC *grpc.Server, auth Auth) {
	ssov1.RegisterAuthServer(gRPC, &serverAPI{
		auth: auth,
	})
}

func (s *serverAPI) Login(ctx context.Context, req *ssov1.LoginRequest) (*ssov1.LoginResponse, error) {
	panic("implement me")
}

func (s *serverAPI) Register(ctx context.Context, req *ssov1.RegisterRequest) (*ssov1.RegisterResponse, error) {
	panic("implement me")
}

func (s *serverAPI) IsAdmin(ctx context.Context, req *ssov1.IsAdminRequest) (*ssov1.IsAdminResponse, error) {
	panic("implement me")
}

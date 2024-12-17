package grpc

import (
	"auth_service/internal/jwt"
	"auth_service/internal/models"
	"auth_service/internal/repositiry/storage"
	//"auth_service/pkg/api"
	"context"
	api "github.com/artemSorokin1/Auth-proto/protos/gen/protos/proto"
	"golang.org/x/crypto/bcrypt"
	"log/slog"
)

const (
	secretKey = "efjodliwjr0o2ijfnoewlkf0p21ASD3"
)

type AuthService struct {
	api.UnimplementedAuthServiceServer
	stor *storage.Storage
}

func (s *AuthService) Register(ctx context.Context, req *api.RegisterRequest) (*api.RegisterResponse, error) {
	slog.Info("Register method called")

	passHash, err := bcrypt.GenerateFromPassword([]byte(req.GetPassword()), bcrypt.DefaultCost)
	if err != nil {
		slog.Warn("error hashing password")
		return nil, err
	}

	user := models.User{
		Email:    req.GetEmail(),
		Username: req.GetUsername(),
		PassHash: string(passHash),
		Role:     "user",
	}

	id, err := s.stor.AddNewUser(user)
	if err != nil || id == -1 {
		slog.Warn("error adding new user")
		return nil, err
	}

	return &api.RegisterResponse{
		UserId: id,
	}, nil

}

func (s *AuthService) Login(ctx context.Context, req *api.LoginRequest) (*api.LoginResponse, error) {
	slog.Info("Login method called")

	user, err := s.stor.GetUser(req.GetUsername())
	if err != nil {
		slog.Warn("user not found in login")
		return nil, storage.ErrUserNotFound
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PassHash), []byte(req.GetPassword())); err != nil {
		slog.Warn("error comparing hashes")
		return nil, storage.ErrUserNotFound
	}

	token, err := jwt.NewToken(&user, secretKey)
	if err != nil {
		return nil, err
	}

	return &api.LoginResponse{
		Token: token,
	}, nil
}

func (s *AuthService) IsAdmin(ctx context.Context, req *api.IsAdminRequest) (*api.IsAdminResponse, error) {
	slog.Info("IsAdmin method called")

	isAdmin, err := s.stor.IsAdmin(req.GetUserId())
	if err != nil {
		return nil, err
	}

	return &api.IsAdminResponse{
		IsAdmin: isAdmin,
	}, nil

}

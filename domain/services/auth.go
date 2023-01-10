package services

import (
	"context"
	"net/http"

	"github.com/INVITATION-RPC-AUTH/domain/models"
	"github.com/INVITATION-RPC-AUTH/domain/repository"
	"github.com/INVITATION-RPC-AUTH/pkg/pb"
	"github.com/INVITATION-RPC-AUTH/pkg/utils"
	"google.golang.org/grpc/status"
)

type AuthService interface {
	Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error)
	Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error)
}

type authService struct {
	AuthRepository repository.AuthRepository
	Jwt            utils.JwtWrapper

	// cache, config, db transaction etc inject here
}

func NewAuthService() *authService {
	return &authService{}
}

func (s *authService) SetAuthRepository(userRepo repository.AuthRepository) *authService {
	s.AuthRepository = userRepo
	return s
}

func (s *authService) Validate() *authService {
	if s.AuthRepository == nil {
		panic("handler need user repository")
	}

	return s
}

func (s *authService) Register(ctx context.Context, user *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	comp := models.User{
		Name:     user.GetName(),
		Email:    user.GetEmail(),
		Password: utils.HashPassword(user.Password),
	}

	if _, err := s.AuthRepository.FindUserByEmail(comp.Email); err == nil {
		return nil, status.Error(400, "Email has been registered")
	}

	if _, err := s.AuthRepository.Insert(comp); err != nil {
		return nil, err
	}

	return &pb.RegisterResponse{
		Status: http.StatusCreated,
	}, nil
}

func (s *authService) Login(ctx context.Context, user *pb.LoginRequest) (*pb.LoginResponse, error) {

	result, err := s.AuthRepository.FindUserByEmail(user.Email)

	if err != nil {
		return nil, status.Error(400, "Email has'n been registered")
	}

	match := utils.CheckPasswordHash(user.Password, result.Password)

	if !match {
		return nil, status.Error(400, "Wrong password")
	}

	token, _ := s.Jwt.GenerateToken(*result)

	return &pb.LoginResponse{
		Status: http.StatusOK,
		Token:  token,
	}, nil
}

package services

import (
	"context"
	"net/http"

	"github.com/INVITATION-RPC-AUTH/domain/models"
	"github.com/INVITATION-RPC-AUTH/domain/redis"
	"github.com/INVITATION-RPC-AUTH/domain/repository"
	"github.com/INVITATION-RPC-AUTH/pkg/pb"
	"github.com/INVITATION-RPC-AUTH/pkg/utils"
	"google.golang.org/grpc/status"
)

type AuthService interface {
	Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error)
	Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error)
	Logout(ctx context.Context, req *pb.LogoutRequest) (*pb.LogoutResponse, error)
	RefreshToken(ctx context.Context, refreshToken *pb.RefreshTokenRequest) (*pb.RefreshTokenResponse, error)
}

type authService struct {
	AuthRepository  repository.AuthRepository
	TokenRepository repository.TokenRepository
	AuthRedis       redis.AuthRedis
	Jwt             utils.JwtWrapper
	// cache, config, db transaction etc inject here
}

func NewAuthService() *authService {
	return &authService{}
}

func (s *authService) SetAuthRepository(userRepo repository.AuthRepository) *authService {
	s.AuthRepository = userRepo
	return s
}

func (s *authService) SetTokenRepository(tokenRepo repository.TokenRepository) *authService {
	s.TokenRepository = tokenRepo
	return s
}

func (s *authService) SetAuthRedis(authRedis redis.AuthRedis) *authService {
	s.AuthRedis = authRedis
	return s
}

func (s *authService) SetAuthJwt(authJwt utils.JwtWrapper) *authService {
	s.Jwt = authJwt
	return s
}

func (s *authService) Validate() *authService {
	if s.AuthRepository == nil {
		panic("handler need user repository")
	}

	if s.AuthRedis == nil {
		panic("handler need user redis")
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

	token, err := s.Jwt.GenerateAccessToken(*result)

	if err != nil {
		return nil, status.Error(400, err.Error())
	}

	s.AuthRedis.SetAuthRedis(ctx, result.Email, token)

	if _, err := s.TokenRepository.Insert(models.Token{Token: token, UserID: int(result.Id)}); err != nil {
		return nil, status.Error(400, "Error while saving token")
	}

	return &pb.LoginResponse{
		Status: http.StatusOK,
		Token:  token,
	}, nil
}

func (s *authService) Logout(ctx context.Context, logout *pb.LogoutRequest) (*pb.LogoutResponse, error) {
	// ambil token
	claims, err := s.Jwt.ValidateToken(logout.Token)
	if err != nil {
		return nil, status.Error(401, err.Error())
	}
	// cek and delete token in redis
	s.AuthRedis.DeleteAuthRedis(ctx, claims.Email)

	// cek and delete token in database
	token, err := s.TokenRepository.GetByUserId(int(claims.Id))
	if err != nil {
		return nil, status.Error(401, err.Error())
	}
	if _, err := s.TokenRepository.Delete(int(token.Id)); err != nil {
		return nil, status.Error(401, err.Error())
	}

	return &pb.LogoutResponse{
		Status: http.StatusOK,
	}, nil
}

func (s *authService) RefreshToken(ctx context.Context, refreshToken *pb.RefreshTokenRequest) (*pb.RefreshTokenResponse, error) {

	claims, err := s.Jwt.ValidateToken(refreshToken.Token)

	if err != nil {
		return nil, status.Error(401, err.Error())
	}

	// check email di redis,
	token, err := s.AuthRedis.GetAuthRedis(ctx, claims.Email)

	if err != nil {
		return nil, status.Error(401, err.Error())
	}

	if token != refreshToken.Token {
		return nil, status.Error(401, "User has been signout")
	}

	// jika tidak maka ambil lagi ke database dengan tokennya

	// var user models.User

	// if result := s.H.DB.Where(&models.User{Email: claims.Email}).First(&user); result.Error != nil {
	// 	return &pb.ValidateResponse{
	// 		Status: http.StatusNotFound,
	// 		Error:  "User not found",
	// 	}, nil
	// }

	return &pb.RefreshTokenResponse{
		Status: http.StatusOK,
		UserId: claims.Id,
	}, nil
}

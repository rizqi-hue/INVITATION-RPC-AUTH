package handler

import (
	"context"
	"net/http"

	"github.com/INVITATION-RPC-AUTH/pkg/pb"
)

func (s Server) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {

	_, err := s.AuthService.Register(ctx, req)
	if err != nil {
		return nil, err
	}

	return &pb.RegisterResponse{
		Status: http.StatusCreated,
	}, nil
}

func (s *Server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {

	result, err := s.AuthService.Login(ctx, req)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *Server) RefreshToken(ctx context.Context, req *pb.RefreshTokenRequest) (*pb.RefreshTokenResponse, error) {

	result, err := s.AuthService.RefreshToken(ctx, req)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *Server) Logout(ctx context.Context, req *pb.LogoutRequest) (*pb.LogoutResponse, error) {

	result, err := s.AuthService.Logout(ctx, req)
	if err != nil {
		return nil, err
	}

	return result, nil
}

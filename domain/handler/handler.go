package handler

import (
	"github.com/INVITATION-RPC-AUTH/domain/services"
	"google.golang.org/grpc"
)

type Server struct {
	grpc        *grpc.Server
	AuthService services.AuthService
}

func InitServer(grpc *grpc.Server, auth services.AuthService) *Server {
	return &Server{grpc, auth}
}

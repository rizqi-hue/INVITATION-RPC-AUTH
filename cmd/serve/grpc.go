package serve

import (
	"context"
	"log"
	"net"

	def "github.com/INVITATION-RPC-AUTH/cmd/config"
	"github.com/INVITATION-RPC-AUTH/domain/handler"
	"github.com/INVITATION-RPC-AUTH/domain/repository"
	"github.com/INVITATION-RPC-AUTH/domain/services"
	"github.com/INVITATION-RPC-AUTH/pkg/config"
	"github.com/INVITATION-RPC-AUTH/pkg/pb"
	"google.golang.org/grpc"
)

func NewGrpcServer(ctx context.Context) error {
	if err := config.Load(def.DefaultConfig, Config); err != nil {
		log.Fatal(err)
	}

	lis, err := net.Listen("tcp", config.GetString("port"))

	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	userRepo := repository.NewAuthRepository()
	userService := services.NewAuthService().
		SetAuthRepository(userRepo).
		Validate()

	s := grpc.NewServer()
	handler := handler.InitServer(s, userService)
	pb.RegisterAuthServiceServer(s, handler)

	log.Println("Serving gRPC on 0.0.0.0:" + config.GetString("port"))
	s.Serve(lis)

	return nil
}

package serve

import (
	"context"
	"log"
	"net"
	"strconv"

	def "github.com/INVITATION-RPC-AUTH/cmd/config"
	"github.com/INVITATION-RPC-AUTH/domain/handler"
	"github.com/INVITATION-RPC-AUTH/domain/redis"
	"github.com/INVITATION-RPC-AUTH/domain/repository"
	"github.com/INVITATION-RPC-AUTH/domain/services"
	"github.com/INVITATION-RPC-AUTH/pkg/config"
	"github.com/INVITATION-RPC-AUTH/pkg/pb"
	"github.com/INVITATION-RPC-AUTH/pkg/utils"
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

	exp_auth, _ := strconv.ParseInt(config.GetString("exp_auth"), 10, 64)
	jwt := utils.JwtWrapper{
		SecretKey:       config.GetString("secret"),
		Issuer:          "go-grpc-auth",
		ExpirationHours: exp_auth * 365,
	}
	
	userRepo := repository.NewAuthRepository()

	tokenRepo := repository.NewTokenRepository()
	userRedis := redis.NewAuthRedis(ctx)
	userService := services.NewAuthService().
		SetAuthRepository(userRepo).
		SetTokenRepository(tokenRepo).
		SetAuthRedis(userRedis).
		SetAuthJwt(jwt).
		Validate()

	s := grpc.NewServer()
	handler := handler.InitServer(s, userService)
	pb.RegisterAuthServiceServer(s, handler)

	log.Println("Serving gRPC on 0.0.0.0:" + config.GetString("port"))
	s.Serve(lis)

	return nil
}

/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package serve

import (
	"context"
	"log"

	def "github.com/INVITATION-RPC-AUTH/cmd/config"
	"github.com/INVITATION-RPC-AUTH/pkg/config"
	"github.com/INVITATION-RPC-AUTH/pkg/database"
	"github.com/INVITATION-RPC-AUTH/pkg/logger"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var (
	Config string

	ServeCmd = &cobra.Command{
		Use:   "serve",
		Short: "configration & connection",
		Run: func(cmd *cobra.Command, args []string) {
			startServer()
		},
	}
)

func startServer() {
	ctx := context.Background()

	if err := config.Load(def.DefaultConfig, Config); err != nil {
		log.Fatal(err)
	}

	logger.Configure()
	database.PostgresConnection(ctx)

	err := NewGrpcServer(ctx)
	if err != nil {
		panic(err)
	}
}

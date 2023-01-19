package database

import (
	"context"

	"github.com/INVITATION-RPC-AUTH/domain/models"
	"github.com/INVITATION-RPC-AUTH/pkg/config"
	"github.com/INVITATION-RPC-AUTH/pkg/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func PostgresConnection(ctx context.Context) {
	l := logger.GetLoggerContext(ctx, "database", "Connect")
	dsn := config.GetString("postgres_dsn")

	// l.Info(dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	l.Info("Connected to postgres")

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Token{})

	DB = db
}

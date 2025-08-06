package db

import (
	"github.com/aquaswim/moyoki/internal/adapter/db/repositories/model"
	"github.com/aquaswim/moyoki/internal/config"
	"github.com/aquaswim/moyoki/internal/core/port"
	"github.com/glebarez/sqlite"
	"github.com/gofiber/fiber/v2/log"
	"gorm.io/gorm"
)

func Connect(cfg *config.DBConfig) *gorm.DB {
	log.Info("connecting to database")
	db, err := gorm.Open(sqlite.Open(cfg.DSN), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %s", err)
	}

	log.Info("connected to database")
	log.Info("Starting migrations")
	err = db.AutoMigrate(&model.Route{})
	if err != nil {
		log.Fatalf("failed to migrate database: %s", err)
	}
	log.Info("Migrations completed")

	return db
}

func Closer(db *gorm.DB) port.Closer {
	return func() error {
		sqlDB, err := db.DB()
		if err != nil {
			return err
		}
		log.Info("closing database connection")
		defer log.Info("database connection closed")
		return sqlDB.Close()
	}
}

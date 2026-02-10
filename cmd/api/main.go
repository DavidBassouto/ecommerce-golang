package main

import (
	"github.com/davidbassouto/ecommerce-golang/internal/config"
	"github.com/davidbassouto/ecommerce-golang/internal/database"
	"github.com/davidbassouto/ecommerce-golang/internal/logger"
	"github.com/gin-gonic/gin"
)

func main() {
	log := logger.New()
	cfg, err := config.Load()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to load config")
	}

	db, err := database.New(cfg.Database)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect to database")
	}

	mainDb, err := db.DB()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to get database connection")
	}
	defer mainDb.Close()
	gin.SetMode(cfg.Server.GinMode)

	log.Info().Msg("starting server")
}

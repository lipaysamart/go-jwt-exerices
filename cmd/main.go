package main

import (
	"log"

	"github.com/lipaysamart/go-jwt-exerices/internal/bootstrap"
	"github.com/lipaysamart/go-jwt-exerices/internal/model"
	"github.com/lipaysamart/go-jwt-exerices/pkg/config"
	"github.com/lipaysamart/gocommon/dbs"
)

func main() {
	cfg := config.LoadConfig()

	database, err := dbs.NewMySQL(cfg.DatabaseURI)
	if err != nil {
		log.Fatal("Failed serve database:", err)
	}
	if err := database.Migrate(&model.User{}); err != nil {
		log.Fatal("Failed to migrate the database:", err)
	}

	serve := bootstrap.NewBootStrap(database)
	if err := serve.Run(); err != nil {
		log.Fatal("Failed to start the server:", err)
	}

}

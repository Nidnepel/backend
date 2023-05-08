package main

import (
	"fmt"
	"github.com/Nidnepel/backend/internal/config"
	"github.com/Nidnepel/backend/internal/database"
	"github.com/Nidnepel/backend/internal/handler"
	"github.com/Nidnepel/backend/internal/repository"
	"github.com/Nidnepel/backend/internal/service"
	"github.com/Nidnepel/backend/server"
	"github.com/pressly/goose/v3"
	"log"
)

const (
	serverPort = "8000"
)

func main() {
	cfg := config.NewConfig()

	db, err := database.New(
		fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
			cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresDbHost, cfg.PostgresDb),
	)
	if err != nil {
		log.Fatalf("невозможно подключиться к базе: %v", err)
	}

	err = goose.Up(db.DB, "./migrations")
	if err != nil {
		log.Fatalf("невозможно накатить миграции: %v", err)
	}
	defer goose.Down(db.DB, "./migrations")

	repos := repository.NewRepository(database.NewPGX(db))
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(server.Server)
	if err := srv.Run(serverPort, handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

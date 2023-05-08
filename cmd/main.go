package main

import (
	"github.com/Nidnepel/backend/internal/handler"
	"github.com/Nidnepel/backend/internal/repository"
	"github.com/Nidnepel/backend/internal/service"
	"github.com/Nidnepel/backend/server"
	"log"
)

const (
	serverPort = "8000"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(server.Server)
	if err := srv.Run(serverPort, handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

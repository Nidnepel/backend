package main

import (
	"github.com/Nidnepel/backend/internal/handler"
	"github.com/Nidnepel/backend/server"
	"log"
)

const (
	serverPort = "8000"
)

func main() {
	handlers := new(handler.Handler)

	srv := new(server.Server)
	if err := srv.Run(serverPort, handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

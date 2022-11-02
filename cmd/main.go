package main

import (
	"github.com/Vzhrkv/avito_internship/internal"
	"github.com/Vzhrkv/avito_internship/pkg/handler"
	"github.com/Vzhrkv/avito_internship/pkg/repository"
	"github.com/Vzhrkv/avito_internship/pkg/service"
	"log"
)

func main() {
	repo := repository.NewRepository()
	service_main := service.NewService(repo)
	handler_main := handler.NewHandler(service_main)

	srv := new(server.Server)
	if err := srv.Run("8000", handler_main.InitRoutes()); err != nil {
		log.Fatal(err)
	}
}

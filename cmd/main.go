package main

import (
	"awesomeProject/internal/api"
	"awesomeProject/internal/services"
	"awesomeProject/internal/storage"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	addr := ":3030"
	postgresDb, err := storage.ConnectDb()

	if err != nil {
		log.Fatal(err)
	}

	taskService := services.NewService(postgresDb)

	app := fiber.New()
	api.TaskApiGroup(app, taskService)
	log.Fatal(app.Listen(addr))

}

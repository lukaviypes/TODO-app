package main

import (
	"awesomeProject/internal/api"
	"awesomeProject/internal/config"
	"awesomeProject/internal/services"
	"awesomeProject/internal/storage"
	"log"
)

func main() {
	cfg, err := config.NewConfig()

	postgresDb, err := storage.ConnectDb(cfg.DB_user, cfg.DB_pass, cfg.DB_name)
	defer postgresDb.CloseDb()
	if err != nil {
		log.Fatal(err)
	}

	taskService := services.NewService(postgresDb, cfg.JWT_secret)
	app := api.NewServer(taskService)

	err = app.Driver.Listen(cfg.Address)
	log.Fatal(err)

}

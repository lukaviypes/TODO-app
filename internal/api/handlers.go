package api

import (
	"awesomeProject/internal/services"
	"github.com/gofiber/fiber/v2"
)

type Server struct {
	Driver  *fiber.App
	Service *services.Service
}

func NewServer(service *services.Service) *Server {
	server := &Server{fiber.New(), service}

	server.Driver.Post("\tasks", server.TaskCreatehandler())

	return server
}

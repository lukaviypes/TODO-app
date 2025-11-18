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
	tasks := server.Driver.Group("/tasks", server.UserAuthMiddleware())
	tasks.Post("/tasks", server.TaskCreatehandler())

	auth := server.Driver.Group("/login")
	auth.Post("/user", server.UserCreatehandler())
	auth.Get("/user", server.UserLoginhandler())
	return server
}

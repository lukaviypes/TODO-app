package api

import (
	"awesomeProject/internal/services"
	"github.com/gofiber/fiber/v2"
)

func TaskApiGroup(a *fiber.App, s services.Service) {

	a.Post("/tasks", TaskCreatehandler(s))

}

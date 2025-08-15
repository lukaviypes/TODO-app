package api

import (
	"awesomeProject/internal/models"
	"awesomeProject/internal/services"
	"awesomeProject/internal/storage"
	"github.com/gofiber/fiber/v2"
)

type Server struct {
	Db   *storage.DataBase
	App  *fiber.App
	Addr string
}

func NewServer(addr string) *Server {
	return &Server{
		Db:   storage.ConnectDb(),
		App:  fiber.New(),
		Addr: addr,
	}
}

func (s *Server) Init() error {
	s.App.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"hello": "world",
		})
	})
	s.App.Post("/", TaskCreatehandler)

	if err := s.App.Listen(":" + s.Addr); err != nil {
		return err
	}
	return nil
}

func TaskCreatehandler(s services.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var reqbody models.Task
		if err := c.BodyParser(&reqbody.Title); err != nil {
			return c.SendStatus(400)
		}
		result, err := s.CreateTask(reqbody.Title)
		if err != nil {
			return c.SendStatus(500)
		}
		return c.
	}

}

package api

import (
	"awesomeProject/internal/models"
	"awesomeProject/internal/services"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

//type Server struct {
//	Db   *storage.DataBase
//	App  *fiber.App
//	Addr string
//}
//
//func NewServer(addr string) *Server {
//	return &Server{
//		Db:   storage.ConnectDb(),
//		App:  fiber.New(),
//		Addr: addr,
//	}
//}
//
//func (s *Server) Init() error {
//	s.App.Get("/", func(c *fiber.Ctx) error {
//		return c.JSON(fiber.Map{
//			"hello": "world",
//		})
//	})
//	s.App.Post("/", TaskCreatehandler)
//
//	if err := s.App.Listen(":" + s.Addr); err != nil {
//		return err
//	}
//	return nil
//}

func TaskCreatehandler(s services.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		reqbody := new(models.Task)
		if err := c.BodyParser(reqbody); err != nil {

			return c.SendStatus(fiber.StatusBadRequest)
		}
		fmt.Println(reqbody)
		result, err := s.CreateTask(reqbody.Title)
		if err != nil {
			return c.SendStatus(500)
		}
		respBody := models.Task{
			Id:    result,
			Title: reqbody.Title,
		}
		return c.JSON(respBody)
	}

}

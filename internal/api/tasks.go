package api

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type CreateTaskReq struct {
	Title string `json:"title"`
}
type CreateTaskResp struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
}

func (s *Server) TaskCreatehandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		reqbody := new(CreateTaskReq)
		if err := c.BodyParser(reqbody); err != nil {

			return c.SendStatus(http.StatusBadRequest)
		}

		result, err := s.Service.CreateTask(reqbody.Title)
		if err != nil {
			return c.SendStatus(http.StatusInternalServerError)
		}
		respBody := CreateTaskResp{
			ID:    result,
			Title: reqbody.Title,
		}
		return c.JSON(respBody)
	}

}

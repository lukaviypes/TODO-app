package api

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type UserReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserResp struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}

func (s *Server) UserCreatehandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		reqbody := new(UserReq)
		if err := c.BodyParser(reqbody); err != nil {

			return c.SendStatus(http.StatusBadRequest)
		}

		result, err := s.Service.CreateTask(reqbody.Title)
		if err != nil {
			return c.SendStatus(http.StatusInternalServerError)
		}

		return c.JSON(respBody)
	}

}

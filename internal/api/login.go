package api

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"
	"strings"
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

		err := s.Service.CreateUser(reqbody.Username, reqbody.Password)
		if err != nil {
			c.SendString(err.Error())
			return c.SendStatus(http.StatusInternalServerError)
		}

		return c.SendStatus(http.StatusOK)
	}

}

func (s *Server) UserLoginhandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		reqbody := new(UserReq)
		if err := c.BodyParser(reqbody); err != nil {
			return c.SendStatus(http.StatusBadRequest)
		}
		stringtoken, err := s.Service.GetToken(reqbody.Username, reqbody.Password)
		if err != nil {
			c.SendString("invalid username or password")
			return c.SendStatus(http.StatusUnauthorized)
		}
		c.SendString(stringtoken)
		return c.SendStatus(http.StatusOK)
	}
}

func (s *Server) UserAuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authString := c.Get("Authorization")
		if authString == "" {
			c.SendString("Authorization required")
			return c.SendStatus(http.StatusUnauthorized)
		}

		autharr := strings.Split(authString, " ")
		if len(autharr) != 2 && autharr[0] != "Bearer" {
			c.SendString("Bearer authorization required")
			return c.SendStatus(http.StatusUnauthorized)
		}
		log.Println(autharr)
		if err := s.Service.ValidateToken(autharr[1]); err != nil {
			c.SendString(err.Error())
			return c.SendStatus(http.StatusUnauthorized)
		}

		return c.Next()
	}
}

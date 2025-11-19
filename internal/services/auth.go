package services

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

type User struct {
	Username string
	HashPass string
}

func (s *Service) CreateUser(username, password string) error {

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	log.Print(string(hash))
	if err != nil {
		return errors.New("cannot hash password")
	}
	user := User{
		Username: username,
		HashPass: string(hash),
	}
	err = s.Repo.CreateUser(user.Username, user.HashPass)
	if err != nil {
		return err
	}
	return nil

}

func (s *Service) GetToken(username, password string) (string, error) {
	hashpass, err := s.Repo.GetUser(username)
	if err != nil {
		return "", err
	}
	err = bcrypt.CompareHashAndPassword([]byte(hashpass), []byte(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 12).Unix(),
		"iat":      time.Now().Unix(),
	})
	StringToken, err := token.SignedString([]byte(s.Secret))
	if err != nil {
		return "", errors.New(fmt.Sprint("error signing token:", err))
	}
	return StringToken, nil
}

func (s *Service) ValidateToken(tokenstring string) error {

	token, err := jwt.Parse(tokenstring, func(token *jwt.Token) (any, error) {
		return []byte(s.Secret), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))

	if err != nil {
		log.Fatal(err)
	}

	log.Println(token)

	if !token.Valid {
		return errors.New("invalid token")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return errors.New("invalid token")
	}
	//username, ok := claims["username"].(string)
	if Expired, ok := claims["exp"].(int64); ok {
		if Expired < time.Now().Unix() {
			return errors.New("token is expired")
		}
		return nil
	}
	return nil
}

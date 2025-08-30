package services

import (
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username string
	HashPass string
}

func (s *Service) CreateUser(username, password string) error {

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return err
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
	token := jwt.New(jwt.SigningMethodHS256)
	StringToken, err := token.SignedString([]byte(s.Secret))
	if err != nil {
		return "", err
	}
	return StringToken, nil
}

//func (s *Service) ValidateToken(token string) error {
//
//	_, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {})
//}

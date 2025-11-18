package config

import (
	"errors"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	Address    string `json:"address"`
	DB_user    string `json:"db_user"`
	DB_pass    string `json:"db_pass"`
	DB_name    string `json:"db_name"`
	JWT_secret string `json:"jwt_secret"`
}

func NewConfig() (*Config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, errors.New("error loading .env file")
	}
	cfg := Config{
		Address:    os.Getenv("ADDRESS"),
		DB_user:    os.Getenv("DB_USER"),
		DB_pass:    os.Getenv("DB_PASS"),
		DB_name:    os.Getenv("DB_NAME"),
		JWT_secret: os.Getenv("JWT_SECRET"),
	}

	return &cfg, nil
}

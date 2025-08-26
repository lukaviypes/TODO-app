package services

import (
	"awesomeProject/internal/storage"
)

type Service interface {
	CreateTask(title string) (int64, error)
}

type service struct {
	repo storage.Storage
}

func NewService(repo storage.Storage) Service {
	return &service{repo: repo}
}

func (s *service) CreateTask(title string) (int64, error) {
	id, err := s.repo.InsertTask(title)

	if err != nil {
		return 0, err
	}

	return id, nil

}

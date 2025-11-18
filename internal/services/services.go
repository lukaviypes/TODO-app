package services

import (
	"awesomeProject/internal/storage"
)

type Service struct {
	Repo   *storage.DataBase
	Secret string
}

//type Task struct {
//	ID   int64
//	Title string
//}

func NewService(repo *storage.DataBase, secret string) *Service {
	return &Service{
		Repo:   repo,
		Secret: secret}
}

func (s *Service) CreateTask(title string) (int64, error) {
	id, err := s.Repo.InsertTask(title)

	if err != nil {
		return 0, err
	}

	return id, nil

}

//func (s *Service) GetTask() []string {
//	id, err := s.Repo.InsertTask(title)
//
//	if err != nil {
//		return 0, err
//	}
//
//	return id, nil
//
//}

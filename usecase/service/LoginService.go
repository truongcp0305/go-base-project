package service

import (
	"go-project/model"
	"go-project/usecase/repository"
)

type LoginService struct {
	repo repository.UserRepository
}

func NewLoginService(r repository.UserRepository) *LoginService {
	return &LoginService{
		repo: r,
	}
}

func (s *LoginService) Login(user model.User) error {
	return nil
}

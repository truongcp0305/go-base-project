package service

import (
	"errors"
	"go-project/library"
	"go-project/model"
	"go-project/usecase/repository"
)

type RegisterService struct {
	repo repository.UserRepository
}

func NewRegisterService(r repository.UserRepository) *RegisterService {
	return &RegisterService{
		repo: r,
	}
}

func (s *RegisterService) CreateAccount(user *model.User) error {
	listUser, err := s.repo.FindUserWithUserName(user.UserName)
	if err != nil {
		return err
	}
	if len(listUser) > 0 {
		return errors.New("UserName already exsit")
	}
	token, err := library.CreateJwt(*user)
	if err != nil {
		return err
	}
	user.Token = token
	user.Password = library.HashString(user.Password)
	err = s.repo.CreateUser(user)
	if err != nil {
		return err
	}
	return nil
}

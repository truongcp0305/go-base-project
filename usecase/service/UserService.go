package service

import (
	"errors"
	"go-project/library"
	"go-project/model"
	"go-project/usecase/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(r repository.UserRepository) *UserService {
	return &UserService{
		repo: r,
	}
}

func (s *UserService) CreateAccount(user *model.User) error {
	listUser, err := s.repo.FindUserByUserName(user.UserName)
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

func (s *UserService) Login(user model.User) error {
	listUser, err := s.repo.FindUser(user.UserName, library.HashString(user.Password))
	if err != nil {
		return err
	}
	if len(listUser) == 0 {
		return errors.New("Incorect user name or password")
	}
	return nil
}

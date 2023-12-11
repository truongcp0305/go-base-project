package repository

import "go-project/model"

type UserRepository interface {
	CreateUser(user *model.User) error
	FindUserByUserName(userName string) ([]model.User, error)
	FindUser(userName string, password string) ([]model.User, error)
}

package repository

import "go-project/model"

type UserRepository interface {
	CreateUser(user *model.User) error
	FindUserWithUserName(userName string) ([]model.User, error)
}

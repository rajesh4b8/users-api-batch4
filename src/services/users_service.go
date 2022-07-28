package services

import (
	"github.com/rajesh4b8/users-api-batch4/src/domain/users"
	"github.com/rajesh4b8/users-api-batch4/src/utils/errors"
)

type UserServiceInterface interface {
	GetUser(userId int) (*users.User, *errors.RestErr)
	CreateUser(u users.User) (*users.User, *errors.RestErr)
}

type userService struct {
}

func NewUserService() UserServiceInterface {
	// return &userService{}
	return new(userService)
}

func (u *userService) GetUser(userId int) (*users.User, *errors.RestErr) {
	user, err := users.GetUSerById(userId)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (us *userService) CreateUser(u users.User) (*users.User, *errors.RestErr) {
	// Do validations on user

	restErr := u.Save()
	if restErr != nil {
		return nil, restErr
	}

	return &u, nil
}

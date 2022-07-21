package services

import (
	"github.com/rajesh4b8/users-api-batch4/src/domain/users"
	"github.com/rajesh4b8/users-api-batch4/src/utils/errors"
)

func GetUser(userId int) (*users.User, *errors.RestErr) {
	user, err := users.GetUSerById(userId)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func CreateUser(u users.User) (*users.User, *errors.RestErr) {
	// Do validations on user

	restErr := u.Save()
	if restErr != nil {
		return nil, restErr
	}

	return &u, nil
}

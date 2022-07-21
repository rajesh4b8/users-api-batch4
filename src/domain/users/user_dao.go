package users

import (
	"fmt"
	"time"

	"github.com/rajesh4b8/users-api-batch4/src/utils/errors"
)

var (
	usersDB = make(map[int]*User)
)

func GetUSerById(id int) (*User, *errors.RestErr) {
	user, ok := usersDB[id]
	if !ok {
		return nil, errors.NewNotFoundError(fmt.Sprintf("User with id %v not found", id))
	}

	return user, nil
}

func (user *User) Save() *errors.RestErr {
	if usersDB[user.Id] != nil {
		return errors.NewBadRequestError(fmt.Sprintf("User already exists with id %v", user.Id))
	}

	user.DateCreated = time.Now().Format("2006-01-02T15:04:05Z")
	usersDB[user.Id] = user

	return nil
}

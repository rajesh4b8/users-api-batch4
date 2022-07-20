package users

import (
	"errors"
	"fmt"

	"github.com/rajesh4b8/bookstore_users-api/domain/users"
)

func GetUSerById(id int) (*User, error) {
	if id == 1 {
		dummyUser := users.User{
			Id:        1,
			FirstName: "Rajesh",
			EmailId:   "dummy_user@test.com",
		}
		return &dummyUser, nil
	}

	return nil, errors.New(fmt.Sprintf("User not found for id %v", id))
}

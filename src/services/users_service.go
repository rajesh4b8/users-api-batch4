package services

import "github.com/rajesh4b8/users-api-batch4/src/domain/users"

func GetUser(userId int) (*users.User, error) {
	user, err := users.GetUSerById(userId)
	if err != nil {
		return nil, error.New("")
	}

	return user, nil
}

func CreateUser(u users.User) (users.User, error) {

	return u, nil
}

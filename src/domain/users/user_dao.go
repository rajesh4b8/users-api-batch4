package users

import (
	"fmt"
	"time"

	"github.com/rajesh4b8/users-api-batch4/src/datasource/postgres/users_db"
	"github.com/rajesh4b8/users-api-batch4/src/logger"
	"github.com/rajesh4b8/users-api-batch4/src/utils/errors"
)

var (
	// temp data storage using map, which will be reset everytime we restart the service.
	usersDB = make(map[int]*User)
	log     = logger.GetLogger()
)

func GetUSerById(id int) (*User, *errors.RestErr) {
	user, ok := usersDB[id]
	if !ok {
		return nil, errors.NewNotFoundError(fmt.Sprintf("User with id %v not found", id))
	}

	return user, nil
}

func (user *User) Save() *errors.RestErr {
	// if usersDB[user.Id] != nil {
	// 	return errors.NewBadRequestError(fmt.Sprintf("User already exists with id %v", user.Id))
	// }

	// user.DateCreated = time.Now().Format("2006-01-02T15:04:05Z")
	// usersDB[user.Id] = user

	stmt, err := users_db.Client.Prepare("insert into users(first_name, last_name, email) values ($1, $2, $3) returning user_id, date_created")
	if err != nil {
		log.Error("Error while prepare stmt " + err.Error())
		return errors.NewInternalServerError("Error while Saving to DB")
	}
	defer stmt.Close()

	var dateCreted time.Time
	if err := stmt.QueryRow(user.FirstName, user.LastName, user.EmailId).Scan(&user.Id, &dateCreted); err != nil {
		log.Error("Error while inerting data " + err.Error())
		return errors.NewInternalServerError("Error while Saving data")
	}

	user.DateCreated = dateCreted.Format("2006-01-02T15:04:05Z")

	return nil
}

package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rajesh4b8/users-api-batch4/src/domain/users"
	"github.com/rajesh4b8/users-api-batch4/src/services"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user users.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println("Request body not valid")
		return
	}

	fmt.Println("Input user", user)
	user.Id = 12
	user.DateCreated = "Today!"

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(user)
}

func ReadUserHandler(w http.ResponseWriter, r *http.Request) {
	user, err := services.GetUser(1)
	if err != nil {
		fmt.Println("Error reading user with id")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(user)
}

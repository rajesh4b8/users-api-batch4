package app

import (
	"github.com/rajesh4b8/users-api-batch4/src/controllers/ping"
	"github.com/rajesh4b8/users-api-batch4/src/controllers/user"
)

func mapUrls() {
	router.HandleFunc("/ping", ping.PingHandler).Methods("GET")

	router.HandleFunc("/users", user.CreateUserHandler).Methods("POST")
	router.HandleFunc("/users/{userId}", user.ReadUserHandler).Methods("GET")
}

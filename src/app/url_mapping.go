package app

import (
	"github.com/rajesh4b8/users-api-batch4/src/controllers/ping"
	"github.com/rajesh4b8/users-api-batch4/src/controllers/user"
	"github.com/rajesh4b8/users-api-batch4/src/services"
)

var (
	pingController = ping.NewPingController()
	userController = user.NewUserController(services.NewUserService())
)

func mapUrls() {
	router.HandleFunc("/ping", pingController.PingHandler).Methods("GET")

	router.HandleFunc("/users", userController.CreateUserHandler).Methods("POST")
	router.HandleFunc("/users/{userId}", userController.ReadUserHandler).Methods("GET")
}

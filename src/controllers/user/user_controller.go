package user

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rajesh4b8/users-api-batch4/src/domain/users"
	"github.com/rajesh4b8/users-api-batch4/src/logger"
	"github.com/rajesh4b8/users-api-batch4/src/services"
	"github.com/rajesh4b8/users-api-batch4/src/utils/errors"
)

var (
	log = logger.GetLogger()
	// userService = services.NewUserService()
)

func NewUserController(userService services.UserServiceInterface) userControllerInterface {
	return &userController{userService}
}

type userControllerInterface interface {
	CreateUserHandler(w http.ResponseWriter, r *http.Request)
	ReadUserHandler(w http.ResponseWriter, r *http.Request)
}

type userController struct {
	userService services.UserServiceInterface
}

func (c *userController) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user users.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Error(fmt.Sprintf("Error parsing the body \n%v", r.Body))
		errors.NewBadRequestError("Request body is not a valid json").HandleError(w)
		return
	}

	u, restErr := c.userService.CreateUser(user)
	if restErr != nil {
		restErr.HandleError(w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(u)
}

func (c *userController) ReadUserHandler(w http.ResponseWriter, r *http.Request) {
	// Getting the parameters
	vars := mux.Vars(r)

	// converting string to int
	userId, err := strconv.Atoi(vars["userId"])
	if err != nil {
		errors.NewBadRequestError("userId must be a number").HandleError(w)
		return
	}

	user, restErr := c.userService.GetUser(userId)
	if restErr != nil {
		restErr.HandleError(w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

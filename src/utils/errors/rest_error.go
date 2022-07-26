package errors

import (
	"encoding/json"
	"net/http"
)

type RestErr struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Error   string `json:"error"`
}

func (err *RestErr) HandleError(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(err.Status)
	json.NewEncoder(w).Encode(err)
}

func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		Status:  http.StatusNotFound,
		Message: message,
		Error:   "Not Found",
	}
}

func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Status:  http.StatusBadRequest,
		Message: message,
		Error:   "Bad Request",
	}
}

func NewInternalServerError(message string) *RestErr {
	return &RestErr{
		Status:  http.StatusInternalServerError,
		Message: message,
		Error:   "Internal Server Error",
	}
}

package app

import (
	"net/http"

	"github.com/gorilla/mux"
)

var (
	router = mux.NewRouter()
)

func StartApplication() {
	mapUrls()

	http.ListenAndServe("127.0.0.1:8080", router)
}

package app

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rajesh4b8/users-api-batch4/src/logger"
)

var (
	router = mux.NewRouter()
	log    = logger.GetLogger()
)

func StartApplication() {
	mapUrls()

	err := http.ListenAndServe("127.0.0.1:8080", router)
	if err != nil {
		log.Fatal("Trouble starting your application " + err.Error())
	}
}

package app

import "github.com/rajesh4b8/users-api-batch4/src/controllers/ping"

func mapUrls() {
	router.HandleFunc("/ping", ping.PingHandler).Methods("GET")
}

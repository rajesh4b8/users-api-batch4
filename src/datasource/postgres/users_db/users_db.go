package users_db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/rajesh4b8/users-api-batch4/src/logger"
)

const (
	username = "postgres"
	password = "dev"
	host     = "127.0.0.1"
	schema   = "postgres"
)

var (
	Client *sql.DB
	log    = logger.GetLogger()
)

func init() {
	datasource := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
		username,
		password,
		host,
		schema,
	)

	var err error
	Client, err = sql.Open("postgres", datasource)

	if err != nil {
		log.Error("Something wrong with DB! " + err.Error())
		panic(err)
	}

	if err := Client.Ping(); err != nil {
		log.Error("DB Connection failed!" + err.Error())
		panic(err)
	}

	log.Info("Database connection successful!")
}

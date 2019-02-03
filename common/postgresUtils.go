package common

import (
	"fmt"
	"log"

	"database/sql"
	_ "github.com/lib/pq"
)

var Db *sql.DB

func params() string {
	info := fmt.Sprintf("host=%s port=%s dbname=%s "+
		"sslmode=%s user=%s password=%s ",
		AppConfig.DBHost,
		AppConfig.DBPort,
		AppConfig.DBName,
		"disable",
		AppConfig.DBUser,
		AppConfig.DBPwd,
	)
	return info
}

func fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func newDbSession() {
	var err error
	Db, err = sql.Open("postgres", params())
	fatal(err)
}

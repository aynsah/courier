package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func LoadDatabase() error {
	var err error

	DB, err = sql.Open("mysql", "root:@tcp(localhost:3306)/courier")
	if err != nil {
		return err
	}

	if err = DB.Ping(); err != nil {
		return err
	}

	return err
}

package database

import (
	"courier/services/user/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func LoadDatabase() error {
	var err error

	DB, err = sql.Open(config.Config.DBType, config.Config.DBUsername+":"+config.Config.DBPassword+"@tcp("+config.Config.DBAddress+")/"+config.Config.DBName)
	if err != nil {
		return err
	}

	if err = DB.Ping(); err != nil {
		return err
	}

	return err
}

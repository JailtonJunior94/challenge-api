package database

import (
	"database/sql"
	"fmt"
	"net/url"

	"github.com/jailtonjunior94/challenge/configs"
)

func NewSqlServerConnection(config *configs.Conf) (*sql.DB, error) {
	query := url.Values{}
	query.Add("database", config.DBName)
	conn := &url.URL{
		Scheme:   "sqlserver",
		User:     url.UserPassword(config.DBUser, config.DBPassword),
		Host:     fmt.Sprintf("%s:%s", config.DBHost, config.DBPort),
		RawQuery: query.Encode(),
	}

	db, err := sql.Open(config.DBDriver, conn.String())
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(db)
	}

	return db, nil
}

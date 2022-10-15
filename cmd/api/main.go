package main

import (
	"github.com/jailtonjunior94/challenge/configs"
	"github.com/jailtonjunior94/challenge/pkg/database"

	_ "github.com/denisenkom/go-mssqldb"
)

func main() {
	config, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	_, err = database.NewSqlServerConnection(config)
	if err != nil {
		panic(err)
	}
}

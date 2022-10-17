package main

import (
	"fmt"

	"github.com/jailtonjunior94/challenge/configs"
	"github.com/jailtonjunior94/challenge/internal/infra/facades"
	"github.com/jailtonjunior94/challenge/internal/infra/repositories"
	"github.com/jailtonjunior94/challenge/internal/usecases"
	"github.com/jailtonjunior94/challenge/pkg/database"

	_ "github.com/denisenkom/go-mssqldb"
)

func main() {
	config, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	db, err := database.NewSqlServerConnection(config)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	planetRepository := repositories.NewPlanetRepository(db)
	starWarsAPI := facades.NewStarWarsFacade(config.StarWarsAPI)
	importUseCase := usecases.NewImportPlanetUseCase(planetRepository, starWarsAPI)

	err = importUseCase.Execute()
	fmt.Println(err)
}

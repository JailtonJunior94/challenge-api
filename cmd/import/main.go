package main

import (
	"github.com/jailtonjunior94/challenge/configs"
	"github.com/jailtonjunior94/challenge/internal/infra/facades"
	"github.com/jailtonjunior94/challenge/internal/infra/repositories"
	"github.com/jailtonjunior94/challenge/internal/usecases"
	"github.com/jailtonjunior94/challenge/pkg/database"
	"github.com/jailtonjunior94/challenge/pkg/logger"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/sirupsen/logrus"
)

func main() {
	config, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	file := logger.NewLoggerFile(config.LogPath)
	defer file.Close()

	db, err := database.NewSqlServerConnection(config)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	planetRepository := repositories.NewPlanetRepository(db)
	starWarsAPI := facades.NewStarWarsFacade()
	importUseCase := usecases.NewImportPlanetUseCase(planetRepository, starWarsAPI, config.StarWarsAPI)

	err = importUseCase.Execute()
	if err != nil {
		logrus.Errorf("Erro ao importar planetas, %v", err)
	}
	logrus.Info("Planetas importados com sucesso")
}

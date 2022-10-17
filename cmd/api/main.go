package main

import (
	"net/http"

	"github.com/jailtonjunior94/challenge/configs"
	"github.com/jailtonjunior94/challenge/internal/infra/repositories"
	"github.com/jailtonjunior94/challenge/internal/usecases"
	"github.com/jailtonjunior94/challenge/pkg/database"
	"github.com/jailtonjunior94/challenge/pkg/logger"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/go-chi/chi/v5"
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
	fetchUseCase := usecases.NewFetchHandler(planetRepository)
	removeUseCase := usecases.NewRemoveHandler(planetRepository)

	router := chi.NewRouter()

	router.Get("/planets", fetchUseCase.GetPlanets)
	router.Get("/planets/{id}", fetchUseCase.GetPlanetByID)
	router.Delete("/planets/{id}", removeUseCase.RemovePlanetByID)

	http.ListenAndServe(":8080", router)
}

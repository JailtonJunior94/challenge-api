package main

import (
	"net/http"

	"github.com/jailtonjunior94/challenge/configs"
	"github.com/jailtonjunior94/challenge/internal/infra/facades"
	"github.com/jailtonjunior94/challenge/internal/infra/handlers"
	"github.com/jailtonjunior94/challenge/internal/infra/repositories"
	"github.com/jailtonjunior94/challenge/internal/usecases"
	"github.com/jailtonjunior94/challenge/pkg/database"
	"github.com/jailtonjunior94/challenge/pkg/logger"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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

	/* Infra */
	starWarsAPI := facades.NewStarWarsFacade()
	planetRepository := repositories.NewPlanetRepository(db)

	/* Use Cases */
	findAllUseCase := usecases.NewFindAllUseCase(planetRepository)
	findByIdUseCase := usecases.NewFindByIDUseCase(planetRepository)
	removeUseCase := usecases.NewRemovePlanetUseCase(planetRepository)
	importUseCase := usecases.NewImportPlanetUseCase(planetRepository, starWarsAPI, config.StarWarsAPI)

	/* Handlers */
	planetHandler := handlers.NewPlanetHandler(removeUseCase, findByIdUseCase, findAllUseCase)

	/* Import data */
	importUseCase.Execute()

	router := chi.NewRouter()
	router.Use(middleware.Heartbeat("/health"))

	router.Get("/planets", planetHandler.GetPlanets)
	router.Get("/planets/{id}", planetHandler.GetPlanetByID)
	router.Delete("/planets/{id}", planetHandler.RemovePlanet)

	http.ListenAndServe(config.ServerPort, router)
}

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

	planetRepository := repositories.NewPlanetRepository(db)
	fetchUseCase := usecases.NewFetchHandler(planetRepository)

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/planets/{id}", fetchUseCase.GetPlanetByID)

	http.ListenAndServe(":8080", router)
}

package usecases

import (
	"errors"
	"testing"

	"github.com/jailtonjunior94/challenge/internal/domain/dtos"
	apiMock "github.com/jailtonjunior94/challenge/internal/infra/facades/mocks"
	"github.com/jailtonjunior94/challenge/internal/infra/repositories/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestImportPlanets(t *testing.T) {
	planets := []dtos.PlanetsOutput{
		{
			Name:    "Tatooine",
			Climate: "arid",
			Terrain: "desert",
			Films:   []string{"https://swapi.dev/api/films/1/", "https://swapi.dev/api/films/3/"},
		},
		{
			Name:    "Alderaan",
			Climate: "temperate",
			Terrain: "grasslands, mountains",
			Films:   []string{"https://swapi.dev/api/films/1/"},
		},
	}

	responsePlanets := dtos.PaginateOutput[dtos.PlanetsOutput]{
		Count:    60,
		Next:     "https://swapi.dev/api/planets/?page=2",
		Previous: nil,
		Results:  planets,
	}

	responsePlanetsTwo := dtos.PaginateOutput[dtos.PlanetsOutput]{
		Count:    60,
		Next:     "",
		Previous: nil,
		Results:  planets,
	}

	responseFilms := dtos.FilmsOutput{
		Title:       "A New Hope",
		EpisodeID:   4,
		Director:    "George Lucas",
		ReleaseDate: "1977-05-25",
	}

	planetRepositoryMock := new(mocks.PlanetRepositoryMock)
	planetRepositoryMock.On("CountPlanets", mock.Anything).Return(0, nil)
	planetRepositoryMock.On("AddFilm", mock.Anything).Return(nil)
	planetRepositoryMock.On("AddPlanet", mock.Anything).Return(nil)

	starWarsFacadeMock := new(apiMock.StarWarsFacadeMock)
	starWarsFacadeMock.On("FetchPlanets", mock.Anything).Return(&responsePlanets, nil).Once()
	starWarsFacadeMock.On("FetchPlanets", mock.Anything).Return(&responsePlanetsTwo, nil)
	starWarsFacadeMock.On("FetchFilms", mock.Anything).Return(&responseFilms, nil)

	useCase := NewImportPlanetUseCase(planetRepositoryMock, starWarsFacadeMock, "")
	err := useCase.Execute()

	assert.Nil(t, err)
	planetRepositoryMock.AssertNumberOfCalls(t, "CountPlanets", 1)
	starWarsFacadeMock.AssertNumberOfCalls(t, "FetchFilms", 6)
	starWarsFacadeMock.AssertNumberOfCalls(t, "FetchPlanets", 2)
}

func TestValidateBaseImported(t *testing.T) {
	planetRepositoryMock := new(mocks.PlanetRepositoryMock)
	planetRepositoryMock.On("CountPlanets", mock.Anything).Return(60, nil)

	starWarsFacadeMock := new(apiMock.StarWarsFacadeMock)

	useCase := NewImportPlanetUseCase(planetRepositoryMock, starWarsFacadeMock, "")
	err := useCase.Execute()

	assert.Nil(t, err)
	planetRepositoryMock.AssertNumberOfCalls(t, "CountPlanets", 1)
}

func TestImportWithErrorFetchPlanets(t *testing.T) {
	planetRepositoryMock := new(mocks.PlanetRepositoryMock)
	planetRepositoryMock.On("CountPlanets", mock.Anything).Return(0, nil)

	starWarsFacadeMock := new(apiMock.StarWarsFacadeMock)
	starWarsFacadeMock.On("FetchPlanets", mock.Anything).Return(nil, errors.New("Não foi possível obtér planetas")).Once()

	useCase := NewImportPlanetUseCase(planetRepositoryMock, starWarsFacadeMock, "")
	err := useCase.Execute()

	assert.NotNil(t, err)
	planetRepositoryMock.AssertNumberOfCalls(t, "CountPlanets", 1)
	starWarsFacadeMock.AssertNumberOfCalls(t, "FetchPlanets", 1)
}

func TestImportWithErrorFetchPlanetsTwiceCalled(t *testing.T) {
	planets := []dtos.PlanetsOutput{
		{
			Name:    "Tatooine",
			Climate: "arid",
			Terrain: "desert",
			Films:   []string{"https://swapi.dev/api/films/1/", "https://swapi.dev/api/films/3/"},
		},
		{
			Name:    "Alderaan",
			Climate: "temperate",
			Terrain: "grasslands, mountains",
			Films:   []string{"https://swapi.dev/api/films/1/"},
		},
	}

	responsePlanets := dtos.PaginateOutput[dtos.PlanetsOutput]{
		Count:    60,
		Next:     "https://swapi.dev/api/planets/?page=2",
		Previous: nil,
		Results:  planets,
	}

	planetRepositoryMock := new(mocks.PlanetRepositoryMock)
	planetRepositoryMock.On("CountPlanets", mock.Anything).Return(0, nil)

	starWarsFacadeMock := new(apiMock.StarWarsFacadeMock)
	starWarsFacadeMock.On("FetchPlanets", mock.Anything).Return(&responsePlanets, nil).Once()
	starWarsFacadeMock.On("FetchPlanets", mock.Anything).Return(nil, errors.New("Não foi possível obtér planetas")).Once()

	useCase := NewImportPlanetUseCase(planetRepositoryMock, starWarsFacadeMock, "")
	err := useCase.Execute()

	assert.NotNil(t, err)
	planetRepositoryMock.AssertNumberOfCalls(t, "CountPlanets", 1)
	starWarsFacadeMock.AssertNumberOfCalls(t, "FetchPlanets", 2)
}

func TestImportWithErrorFetchFilms(t *testing.T) {
	planets := []dtos.PlanetsOutput{
		{
			Name:    "Tatooine",
			Climate: "arid",
			Terrain: "desert",
			Films:   []string{"https://swapi.dev/api/films/1/", "https://swapi.dev/api/films/3/"},
		},
		{
			Name:    "Alderaan",
			Climate: "temperate",
			Terrain: "grasslands, mountains",
			Films:   []string{"https://swapi.dev/api/films/1/"},
		},
	}

	responsePlanets := dtos.PaginateOutput[dtos.PlanetsOutput]{
		Count:    60,
		Next:     "https://swapi.dev/api/planets/?page=2",
		Previous: nil,
		Results:  planets,
	}

	responsePlanetsTwo := dtos.PaginateOutput[dtos.PlanetsOutput]{
		Count:    60,
		Next:     "",
		Previous: nil,
		Results:  planets,
	}

	planetRepositoryMock := new(mocks.PlanetRepositoryMock)
	planetRepositoryMock.On("CountPlanets", mock.Anything).Return(0, nil)
	planetRepositoryMock.On("AddFilm", mock.Anything).Return(nil)
	planetRepositoryMock.On("AddPlanet", mock.Anything).Return(nil)

	starWarsFacadeMock := new(apiMock.StarWarsFacadeMock)
	starWarsFacadeMock.On("FetchPlanets", mock.Anything).Return(&responsePlanets, nil).Once()
	starWarsFacadeMock.On("FetchPlanets", mock.Anything).Return(&responsePlanetsTwo, nil)
	starWarsFacadeMock.On("FetchFilms", mock.Anything).Return(nil, errors.New("Não foi possível obtér films"))

	useCase := NewImportPlanetUseCase(planetRepositoryMock, starWarsFacadeMock, "")
	err := useCase.Execute()

	assert.Nil(t, err)
	planetRepositoryMock.AssertNumberOfCalls(t, "CountPlanets", 1)
	starWarsFacadeMock.AssertNumberOfCalls(t, "FetchPlanets", 2)
}

func TestImportWithPlanetEntityInvalid(t *testing.T) {
	planets := []dtos.PlanetsOutput{
		{
			Name:    "",
			Climate: "arid",
			Terrain: "desert",
			Films:   []string{"https://swapi.dev/api/films/1/"},
		},
	}

	responsePlanets := dtos.PaginateOutput[dtos.PlanetsOutput]{
		Count:    60,
		Next:     "https://swapi.dev/api/planets/?page=2",
		Previous: nil,
		Results:  planets,
	}

	responsePlanetsTwo := dtos.PaginateOutput[dtos.PlanetsOutput]{
		Count:    60,
		Next:     "",
		Previous: nil,
		Results:  planets,
	}

	responseFilms := dtos.FilmsOutput{
		Title:       "A New Hope",
		EpisodeID:   4,
		Director:    "George Lucas",
		ReleaseDate: "1977-05-25",
	}

	planetRepositoryMock := new(mocks.PlanetRepositoryMock)
	planetRepositoryMock.On("CountPlanets", mock.Anything).Return(0, nil)
	planetRepositoryMock.On("AddFilm", mock.Anything).Return(nil)
	planetRepositoryMock.On("AddPlanet", mock.Anything).Return(nil)

	starWarsFacadeMock := new(apiMock.StarWarsFacadeMock)
	starWarsFacadeMock.On("FetchPlanets", mock.Anything).Return(&responsePlanets, nil).Once()
	starWarsFacadeMock.On("FetchPlanets", mock.Anything).Return(&responsePlanetsTwo, nil)
	starWarsFacadeMock.On("FetchFilms", mock.Anything).Return(&responseFilms, nil)

	useCase := NewImportPlanetUseCase(planetRepositoryMock, starWarsFacadeMock, "")
	err := useCase.Execute()

	assert.Nil(t, err)
	planetRepositoryMock.AssertNumberOfCalls(t, "CountPlanets", 1)
	starWarsFacadeMock.AssertNumberOfCalls(t, "FetchPlanets", 2)
}

func TestImportWithFilmEntityInvalid(t *testing.T) {
	planets := []dtos.PlanetsOutput{
		{
			Name:    "Tatooine",
			Climate: "arid",
			Terrain: "desert",
			Films:   []string{"https://swapi.dev/api/films/1/"},
		},
	}

	responsePlanets := dtos.PaginateOutput[dtos.PlanetsOutput]{
		Count:    60,
		Next:     "https://swapi.dev/api/planets/?page=2",
		Previous: nil,
		Results:  planets,
	}

	responsePlanetsTwo := dtos.PaginateOutput[dtos.PlanetsOutput]{
		Count:    60,
		Next:     "",
		Previous: nil,
		Results:  planets,
	}

	responseFilms := dtos.FilmsOutput{
		Title:       "",
		EpisodeID:   4,
		Director:    "George Lucas",
		ReleaseDate: "1977-05-25",
	}

	planetRepositoryMock := new(mocks.PlanetRepositoryMock)
	planetRepositoryMock.On("CountPlanets", mock.Anything).Return(0, nil)
	planetRepositoryMock.On("AddFilm", mock.Anything).Return(nil)
	planetRepositoryMock.On("AddPlanet", mock.Anything).Return(nil)

	starWarsFacadeMock := new(apiMock.StarWarsFacadeMock)
	starWarsFacadeMock.On("FetchPlanets", mock.Anything).Return(&responsePlanets, nil).Once()
	starWarsFacadeMock.On("FetchPlanets", mock.Anything).Return(&responsePlanetsTwo, nil)
	starWarsFacadeMock.On("FetchFilms", mock.Anything).Return(&responseFilms, nil)

	useCase := NewImportPlanetUseCase(planetRepositoryMock, starWarsFacadeMock, "")
	err := useCase.Execute()

	assert.Nil(t, err)
	planetRepositoryMock.AssertNumberOfCalls(t, "CountPlanets", 1)
	starWarsFacadeMock.AssertNumberOfCalls(t, "FetchPlanets", 2)
}

func TestImportWithAddPlanetError(t *testing.T) {
	planets := []dtos.PlanetsOutput{
		{
			Name:    "Tatooine",
			Climate: "arid",
			Terrain: "desert",
			Films:   []string{"https://swapi.dev/api/films/1/"},
		},
	}

	responsePlanets := dtos.PaginateOutput[dtos.PlanetsOutput]{
		Count:    60,
		Next:     "https://swapi.dev/api/planets/?page=2",
		Previous: nil,
		Results:  planets,
	}

	responsePlanetsTwo := dtos.PaginateOutput[dtos.PlanetsOutput]{
		Count:    60,
		Next:     "",
		Previous: nil,
		Results:  planets,
	}

	responseFilms := dtos.FilmsOutput{
		Title:       "A New Hope",
		EpisodeID:   4,
		Director:    "George Lucas",
		ReleaseDate: "1977-05-25",
	}

	planetRepositoryMock := new(mocks.PlanetRepositoryMock)
	planetRepositoryMock.On("CountPlanets", mock.Anything).Return(0, nil)
	planetRepositoryMock.On("AddFilm", mock.Anything).Return(nil)
	planetRepositoryMock.On("AddPlanet", mock.Anything).Return(errors.New("SQL Exception"))

	starWarsFacadeMock := new(apiMock.StarWarsFacadeMock)
	starWarsFacadeMock.On("FetchPlanets", mock.Anything).Return(&responsePlanets, nil).Once()
	starWarsFacadeMock.On("FetchPlanets", mock.Anything).Return(&responsePlanetsTwo, nil)
	starWarsFacadeMock.On("FetchFilms", mock.Anything).Return(&responseFilms, nil)

	useCase := NewImportPlanetUseCase(planetRepositoryMock, starWarsFacadeMock, "")
	err := useCase.Execute()

	assert.Nil(t, err)
	planetRepositoryMock.AssertNumberOfCalls(t, "CountPlanets", 1)
	starWarsFacadeMock.AssertNumberOfCalls(t, "FetchPlanets", 2)
}

func TestImportWithAddFilmError(t *testing.T) {
	planets := []dtos.PlanetsOutput{
		{
			Name:    "Tatooine",
			Climate: "arid",
			Terrain: "desert",
			Films:   []string{"https://swapi.dev/api/films/1/"},
		},
	}

	responsePlanets := dtos.PaginateOutput[dtos.PlanetsOutput]{
		Count:    60,
		Next:     "https://swapi.dev/api/planets/?page=2",
		Previous: nil,
		Results:  planets,
	}

	responsePlanetsTwo := dtos.PaginateOutput[dtos.PlanetsOutput]{
		Count:    60,
		Next:     "",
		Previous: nil,
		Results:  planets,
	}

	responseFilms := dtos.FilmsOutput{
		Title:       "A New Hope",
		EpisodeID:   4,
		Director:    "George Lucas",
		ReleaseDate: "1977-05-25",
	}

	planetRepositoryMock := new(mocks.PlanetRepositoryMock)
	planetRepositoryMock.On("CountPlanets", mock.Anything).Return(0, nil)
	planetRepositoryMock.On("AddFilm", mock.Anything).Return(errors.New("SQL Exception"))
	planetRepositoryMock.On("AddPlanet", mock.Anything).Return(nil)

	starWarsFacadeMock := new(apiMock.StarWarsFacadeMock)
	starWarsFacadeMock.On("FetchPlanets", mock.Anything).Return(&responsePlanets, nil).Once()
	starWarsFacadeMock.On("FetchPlanets", mock.Anything).Return(&responsePlanetsTwo, nil)
	starWarsFacadeMock.On("FetchFilms", mock.Anything).Return(&responseFilms, nil)

	useCase := NewImportPlanetUseCase(planetRepositoryMock, starWarsFacadeMock, "")
	err := useCase.Execute()

	assert.Nil(t, err)
	planetRepositoryMock.AssertNumberOfCalls(t, "CountPlanets", 1)
	starWarsFacadeMock.AssertNumberOfCalls(t, "FetchPlanets", 2)
}

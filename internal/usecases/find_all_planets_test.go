package usecases

import (
	"errors"
	"net/http"
	"testing"

	"github.com/jailtonjunior94/challenge/internal/domain/dtos"
	"github.com/jailtonjunior94/challenge/internal/domain/entities"
	"github.com/jailtonjunior94/challenge/internal/infra/repositories/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFindAll(t *testing.T) {
	planet, _ := entities.NewPlanet("Yavin IV", "temperate, tropical", "jungle, rainforests")
	planetTwo, _ := entities.NewPlanet("Yavin IV", "temperate, tropical", "jungle, rainforests")

	planets := []entities.Planet{*planet, *planetTwo}

	planetRepositoryMock := new(mocks.PlanetRepositoryMock)
	planetRepositoryMock.On("FindAll", mock.Anything).Return(2, planets, nil)

	useCase := NewFindAllUseCase(planetRepositoryMock)
	input := dtos.NewFilterPlanetInput("Yavin IV", "1", "10")
	response := useCase.Execute(input)

	assert.NotNil(t, response)
	assert.Equal(t, http.StatusOK, response.StatusCode)
	planetRepositoryMock.AssertNumberOfCalls(t, "FindAll", 1)
}

func TestFindAllSqlNoRows(t *testing.T) {
	planetRepositoryMock := new(mocks.PlanetRepositoryMock)
	planetRepositoryMock.On("FindAll", mock.Anything).Return(0, nil, nil)

	useCase := NewFindAllUseCase(planetRepositoryMock)
	input := dtos.NewFilterPlanetInput("Yavin IV", "1", "10")
	response := useCase.Execute(input)

	assert.NotNil(t, response)
	assert.Equal(t, http.StatusNoContent, response.StatusCode)
	planetRepositoryMock.AssertNumberOfCalls(t, "FindAll", 1)
}

func TestFindAllPlanetWithSQLException(t *testing.T) {
	planetRepositoryMock := new(mocks.PlanetRepositoryMock)
	planetRepositoryMock.On("FindAll", mock.Anything).Return(0, nil, errors.New("SQL EXCEPTION"))

	useCase := NewFindAllUseCase(planetRepositoryMock)
	input := dtos.NewFilterPlanetInput("Yavin IV", "1", "10")
	response := useCase.Execute(input)

	assert.NotNil(t, response)
	assert.Equal(t, http.StatusInternalServerError, response.StatusCode)
	planetRepositoryMock.AssertNumberOfCalls(t, "FindAll", 1)
}

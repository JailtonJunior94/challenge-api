package usecases

import (
	"database/sql"
	"errors"
	"net/http"
	"testing"

	"github.com/jailtonjunior94/challenge/internal/domain/entities"
	"github.com/jailtonjunior94/challenge/internal/infra/repositories/mocks"
	"github.com/jailtonjunior94/challenge/pkg/entity"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFindPlanetById(t *testing.T) {
	planet, _ := entities.NewPlanet("Yavin IV", "temperate, tropical", "jungle, rainforests")

	planetRepositoryMock := new(mocks.PlanetRepositoryMock)
	planetRepositoryMock.On("FindByID", mock.Anything).Return(planet, nil)

	useCase := NewFindByIDUseCase(planetRepositoryMock)

	response := useCase.Execute(entity.NewID().String())
	assert.NotNil(t, response)
	assert.Equal(t, http.StatusOK, response.StatusCode)
	planetRepositoryMock.AssertNumberOfCalls(t, "FindByID", 1)
}

func TestFindByIDWithSqlNoRows(t *testing.T) {
	planetRepositoryMock := new(mocks.PlanetRepositoryMock)
	planetRepositoryMock.On("FindByID", mock.Anything).Return(nil, sql.ErrNoRows)

	useCase := NewFindByIDUseCase(planetRepositoryMock)

	response := useCase.Execute(entity.NewID().String())
	assert.NotNil(t, response)
	assert.Equal(t, http.StatusNoContent, response.StatusCode)
	planetRepositoryMock.AssertNumberOfCalls(t, "FindByID", 1)
}

func TestFindByIDPlanetWithSQLException(t *testing.T) {
	planetRepositoryMock := new(mocks.PlanetRepositoryMock)
	planetRepositoryMock.On("FindByID", mock.Anything).Return(nil, errors.New("SQL Exception"))

	useCase := NewFindByIDUseCase(planetRepositoryMock)

	response := useCase.Execute(entity.NewID().String())
	assert.NotNil(t, response)
	assert.Equal(t, http.StatusInternalServerError, response.StatusCode)
	planetRepositoryMock.AssertNumberOfCalls(t, "FindByID", 1)
}

package usecases

import (
	"database/sql"
	"errors"
	"net/http"
	"testing"

	"github.com/jailtonjunior94/challenge/internal/infra/repositories/mocks"
	"github.com/jailtonjunior94/challenge/pkg/entity"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRemovePlanet(t *testing.T) {
	planetRepositoryMock := new(mocks.PlanetRepositoryMock)
	planetRepositoryMock.On("Remove", mock.Anything).Return(nil)

	useCase := NewRemovePlanetUseCase(planetRepositoryMock)

	response := useCase.Execute(entity.NewID().String())

	assert.NotNil(t, response)
	assert.Equal(t, http.StatusNoContent, response.StatusCode)
	planetRepositoryMock.AssertNumberOfCalls(t, "Remove", 1)
}

func TestRemovePlanetWithSqlNoRows(t *testing.T) {
	planetRepositoryMock := new(mocks.PlanetRepositoryMock)
	planetRepositoryMock.On("Remove", mock.Anything).Return(sql.ErrNoRows)

	useCase := NewRemovePlanetUseCase(planetRepositoryMock)

	response := useCase.Execute(entity.NewID().String())

	assert.NotNil(t, response)
	assert.Equal(t, http.StatusNoContent, response.StatusCode)
	planetRepositoryMock.AssertNumberOfCalls(t, "Remove", 1)
}

func TestRemovePlanetWithSQLException(t *testing.T) {
	planetRepositoryMock := new(mocks.PlanetRepositoryMock)
	planetRepositoryMock.On("Remove", mock.Anything).Return(errors.New("SQL Exception"))

	useCase := NewRemovePlanetUseCase(planetRepositoryMock)

	response := useCase.Execute(entity.NewID().String())

	assert.NotNil(t, response)
	assert.Equal(t, http.StatusInternalServerError, response.StatusCode)
	planetRepositoryMock.AssertNumberOfCalls(t, "Remove", 1)
}

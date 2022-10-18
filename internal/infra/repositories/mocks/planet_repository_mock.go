package mocks

import (
	"github.com/jailtonjunior94/challenge/internal/domain/dtos"
	"github.com/jailtonjunior94/challenge/internal/domain/entities"

	"github.com/stretchr/testify/mock"
)

type PlanetRepositoryMock struct {
	mock.Mock
}

func (r *PlanetRepositoryMock) Remove(id string) error {
	args := r.Called(id)
	return args.Error(0)
}

func (r *PlanetRepositoryMock) AddFilm(f *entities.Film) error {
	args := r.Called(f)
	return args.Error(0)
}

func (r *PlanetRepositoryMock) AddPlanet(p *entities.Planet) error {
	args := r.Called(p)
	return args.Error(0)
}

func (r *PlanetRepositoryMock) FindByID(id string) (*entities.Planet, error) {
	args := r.Called(id)
	result, _ := args.Get(0).(*entities.Planet)
	return result, args.Error(1)
}

func (r *PlanetRepositoryMock) FindAll(f *dtos.FilterPlanetInput) (int, []entities.Planet, error) {
	args := r.Called(f)
	count, _ := args.Get(0).(int)
	result, _ := args.Get(1).([]entities.Planet)
	return count, result, args.Error(2)
}

func (r *PlanetRepositoryMock) CountPlanets(name string) (int, error) {
	args := r.Called(name)
	count, _ := args.Get(0).(int)
	return count, args.Error(1)
}

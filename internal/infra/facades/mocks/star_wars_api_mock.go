package mocks

import (
	"github.com/jailtonjunior94/challenge/internal/domain/dtos"

	"github.com/stretchr/testify/mock"
)

type StarWarsFacadeMock struct {
	mock.Mock
}

func (r *StarWarsFacadeMock) FetchPlanets(uri string) (*dtos.PaginateOutput[dtos.PlanetsOutput], error) {
	args := r.Called(uri)
	result, _ := args.Get(0).(*dtos.PaginateOutput[dtos.PlanetsOutput])
	return result, args.Error(1)
}

func (r *StarWarsFacadeMock) FetchFilms(uri string) (*dtos.FilmsOutput, error) {
	args := r.Called(uri)
	result, _ := args.Get(0).(*dtos.FilmsOutput)
	return result, args.Error(1)
}

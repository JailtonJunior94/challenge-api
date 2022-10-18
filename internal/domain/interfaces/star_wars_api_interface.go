package interfaces

import "github.com/jailtonjunior94/challenge/internal/domain/dtos"

type StarWarsFacade interface {
	FetchFilms(uri string) (*dtos.FilmsOutput, error)
	FetchPlanets(uri string) (*dtos.PaginateOutput[dtos.PlanetsOutput], error)
}

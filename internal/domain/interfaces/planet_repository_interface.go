package interfaces

import (
	"github.com/jailtonjunior94/challenge/internal/domain/dtos"
	"github.com/jailtonjunior94/challenge/internal/domain/entities"
)

type PlanetRepository interface {
	Remove(id string) error
	AddFilm(f *entities.Film) error
	AddPlanet(p *entities.Planet) error
	CountPlanets(name string) (int, error)
	FindByID(id string) (*entities.Planet, error)
	FindAll(f *dtos.FilterPlanetInput) (int, []entities.Planet, error)
}

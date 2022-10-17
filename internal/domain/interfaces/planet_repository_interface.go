package interfaces

import "github.com/jailtonjunior94/challenge/internal/domain/entities"

type PlanetRepository interface {
	FindByID(id string) (*entities.Planet, error)
	FindAll(name string, page int, limit int) ([]entities.Planet, error)
	Remove(id string) error
	AddPlanet(p *entities.Planet) (*entities.Planet, error)
	AddFilm(f *entities.Film) (*entities.Planet, error)
}

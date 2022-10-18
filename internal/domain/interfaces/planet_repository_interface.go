package interfaces

import "github.com/jailtonjunior94/challenge/internal/domain/entities"

type PlanetRepository interface {
	Remove(id string) error
	AddFilm(f *entities.Film) error
	AddPlanet(p *entities.Planet) error
	FindByID(id string) (*entities.Planet, error)
	FindAll(name string, page int, limit int) (int, []entities.Planet, error)
}

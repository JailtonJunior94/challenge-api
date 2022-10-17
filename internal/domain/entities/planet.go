package entities

import (
	"errors"

	"github.com/jailtonjunior94/challenge/pkg/entity"
)

var (
	ErrNameIsRequired    = errors.New("name is required")
	ErrClimateIsRequired = errors.New("climate is required")
	ErrTerrainIsRequired = errors.New("terrain is required")
	ErrFilmIsRequired    = errors.New("film is required")
)

type Planet struct {
	ID      entity.ID `json:"id"`
	Name    string    `json:"name"`
	Climate string    `json:"climate"`
	Terrain string    `json:"terrain"`
	Films   []Film    `json:"films"`
}

func NewPlanet(name, climate, terrain string) (*Planet, error) {
	planet := &Planet{
		ID:      entity.NewID(),
		Name:    name,
		Climate: climate,
		Terrain: terrain,
	}

	err := planet.Validate()
	if err != nil {
		return nil, err
	}

	return planet, nil
}

func (p *Planet) Validate() error {
	if p.Name == "" {
		return ErrNameIsRequired
	}

	if p.Climate == "" {
		return ErrClimateIsRequired
	}

	if p.Terrain == "" {
		return ErrTerrainIsRequired
	}
	return nil
}

func (p *Planet) AddFilm(film *Film) error {
	if film == nil {
		return ErrFilmIsRequired
	}

	p.Films = append(p.Films, *film)
	return nil
}

func (p *Planet) AddFilms(films []Film) error {
	p.Films = films
	return nil
}

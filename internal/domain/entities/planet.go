package entities

import "errors"

var (
	ErrNameIsRequired    = errors.New("name is required")
	ErrClimateIsRequired = errors.New("climate is required")
	ErrTerrainIsRequired = errors.New("terrain is required")
)

type Planet struct {
	Name    string
	Climate string
	Terrain string
	Films   []Films
}

func NewPlanet(name, climate, terrain string) (*Planet, error) {
	planet := &Planet{
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

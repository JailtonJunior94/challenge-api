package entities

import (
	"errors"

	"github.com/jailtonjunior94/challenge/pkg/entity"
)

var (
	ErrTitleIsRequired       = errors.New("title is required")
	ErrDirectorIsRequired    = errors.New("director is required")
	ErrReleaseDateIsRequired = errors.New("release date is required")
	ErrInvalidID             = errors.New("invalid id")
)

type Film struct {
	ID          string `json:"id,omitempty"`
	PlanetID    string `json:"planetId,omitempty"`
	Title       string `json:"title,omitempty"`
	Director    string `json:"director,omitempty"`
	ReleaseDate string `json:"releaseDate,omitempty"`
}

func NewFilm(planetID, title, director, releaseDate string) (*Film, error) {
	Film := &Film{
		ID:          entity.NewID().String(),
		PlanetID:    planetID,
		Title:       title,
		Director:    director,
		ReleaseDate: releaseDate,
	}

	err := Film.Validate()
	if err != nil {
		return nil, err
	}

	return Film, nil
}

func (f *Film) Validate() error {
	if f.Title == "" {
		return ErrTitleIsRequired
	}

	if f.Director == "" {
		return ErrDirectorIsRequired
	}

	if f.ReleaseDate == "" {
		return ErrReleaseDateIsRequired
	}
	return nil
}

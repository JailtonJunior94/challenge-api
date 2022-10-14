package entities

import "errors"

var (
	ErrTitleIsRequired       = errors.New("title is required")
	ErrDirectorIsRequired    = errors.New("director is required")
	ErrReleaseDateIsRequired = errors.New("release date is required")
)

type Films struct {
	Title       string
	Director    string
	ReleaseDate string
}

func NewFilms(title, director, releaseDate string) (*Films, error) {
	Films := &Films{
		Title:       title,
		Director:    director,
		ReleaseDate: releaseDate,
	}

	err := Films.Validate()
	if err != nil {
		return nil, err
	}

	return Films, nil
}

func (f *Films) Validate() error {
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

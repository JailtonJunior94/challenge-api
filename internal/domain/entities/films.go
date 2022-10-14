package entities

import "errors"

var (
	ErrTitleIsRequired       = errors.New("title is required")
	ErrDirectorIsRequired    = errors.New("director is required")
	ErrReleaseDateIsRequired = errors.New("release date is required")
)

type Film struct {
	Title       string
	Director    string
	ReleaseDate string
}

func NewFilm(title, director, releaseDate string) (*Film, error) {
	Film := &Film{
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

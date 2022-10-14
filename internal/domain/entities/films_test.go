package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilm(t *testing.T) {
	testCases := []struct {
		name     string
		input    Film
		expected func(p *Film, err error)
	}{
		{
			name:  "Should created film with success",
			input: Film{Title: "A New Hope", Director: "George Lucas", ReleaseDate: "1977-05-25"},
			expected: func(p *Film, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, p)
				assert.Equal(t, "A New Hope", p.Title)
				assert.Equal(t, "George Lucas", p.Director)
				assert.Equal(t, "1977-05-25", p.ReleaseDate)
				assert.Nil(t, p.Validate())
			},
		},
		{
			name:  "Should validated when title is required",
			input: Film{Title: "", Director: "George Lucas", ReleaseDate: "1977-05-25"},
			expected: func(p *Film, err error) {
				assert.Nil(t, p)
				assert.Equal(t, ErrTitleIsRequired, err)
			},
		},
		{
			name:  "Should validated when director is required",
			input: Film{Title: "A New Hope", Director: "", ReleaseDate: "1977-05-25"},
			expected: func(p *Film, err error) {
				assert.Nil(t, p)
				assert.Equal(t, ErrDirectorIsRequired, err)
			},
		},
		{
			name:  "Should validated when release date is required",
			input: Film{Title: "A New Hope", Director: "George Lucas", ReleaseDate: ""},
			expected: func(p *Film, err error) {
				assert.Nil(t, p)
				assert.Equal(t, ErrReleaseDateIsRequired, err)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			film, err := NewFilm(tc.input.Title, tc.input.Director, tc.input.ReleaseDate)
			tc.expected(film, err)
		})
	}
}

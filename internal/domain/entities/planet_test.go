package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlanet(t *testing.T) {
	testCases := []struct {
		name     string
		input    Planet
		expected func(p *Planet, err error)
	}{
		{
			name:  "Should created Planet with success",
			input: Planet{Name: "Tatooine", Climate: "arid", Terrain: "desert"},
			expected: func(p *Planet, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, p)
				assert.Equal(t, "Tatooine", p.Name)
				assert.Equal(t, "arid", p.Climate)
				assert.Equal(t, "desert", p.Terrain)
				assert.Nil(t, p.Validate())
			},
		},
		{
			name:  "Should validated when name is required",
			input: Planet{Name: "", Climate: "arid", Terrain: "desert"},
			expected: func(p *Planet, err error) {
				assert.Nil(t, p)
				assert.Equal(t, ErrNameIsRequired, err)
			},
		},
		{
			name:  "Should validated when climate is required",
			input: Planet{Name: "Tatooine", Climate: "", Terrain: "desert"},
			expected: func(p *Planet, err error) {
				assert.Nil(t, p)
				assert.Equal(t, ErrClimateIsRequired, err)
			},
		},
		{
			name:  "Should validated when terrain is required",
			input: Planet{Name: "Tatooine", Climate: "arid", Terrain: ""},
			expected: func(p *Planet, err error) {
				assert.Nil(t, p)
				assert.Equal(t, ErrTerrainIsRequired, err)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			planet, err := NewPlanet(tc.input.Name, tc.input.Climate, tc.input.Terrain)
			tc.expected(planet, err)
		})
	}
}

func TestAddFilm(t *testing.T) {
	testCases := []struct {
		name     string
		planet   Planet
		film     *Film
		expected func(p *Planet, err error)
	}{
		{
			name:   "Should add film with success",
			planet: Planet{Name: "Tatooine", Climate: "arid", Terrain: "desert"},
			film:   &Film{Title: "A New Hope", Director: "George Lucas", ReleaseDate: "1977-05-25"},
			expected: func(p *Planet, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, p)
			},
		},
		{
			name:   "Should validated when film is required",
			planet: Planet{Name: "Tatooine", Climate: "arid", Terrain: "desert"},
			film:   nil,
			expected: func(p *Planet, err error) {
				assert.NotNil(t, err)
				assert.Equal(t, ErrFilmIsRequired, err)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			planet, _ := NewPlanet(tc.planet.Name, tc.planet.Climate, tc.planet.Terrain)
			err := planet.AddFilm(tc.film)
			tc.expected(planet, err)
		})
	}
}

func TestAddFilms(t *testing.T) {
	films := []Film{
		{
			Title: "A New Hope", Director: "George Lucas", ReleaseDate: "1977-05-25",
		},
	}

	testCases := []struct {
		name     string
		planet   Planet
		film     []Film
		expected func(p *Planet, err error)
	}{
		{
			name:   "Should add film with success",
			planet: Planet{Name: "Tatooine", Climate: "arid", Terrain: "desert"},
			film:   films,
			expected: func(p *Planet, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, p)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			planet, _ := NewPlanet(tc.planet.Name, tc.planet.Climate, tc.planet.Terrain)
			err := planet.AddFilms(tc.film)
			tc.expected(planet, err)
		})
	}
}

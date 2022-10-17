package usecases

import (
	"fmt"

	"github.com/jailtonjunior94/challenge/internal/domain/dtos"
	"github.com/jailtonjunior94/challenge/internal/domain/entities"
	"github.com/jailtonjunior94/challenge/internal/domain/interfaces"
	"github.com/jailtonjunior94/challenge/internal/infra/facades"
)

type importPlanetUseCase struct {
	PlanetRepository interfaces.PlanetRepository
	StarWarsAPI      *facades.StarWarsFacade
}

func NewImportPlanetUseCase(planetRepository interfaces.PlanetRepository, starWarsAPI *facades.StarWarsFacade) *importPlanetUseCase {
	return &importPlanetUseCase{
		PlanetRepository: planetRepository,
		StarWarsAPI:      starWarsAPI,
	}
}

func (u *importPlanetUseCase) Execute() error {
	p, err := u.StarWarsAPI.FetchPlanets("https://swapi.dev/api/planets?page=1")
	if err != nil {
		panic(err)
	}

	var planets []dtos.PlanetsOutput
	planets = append(planets, p.Results...)

	for p.Next != "" {
		p, err = u.StarWarsAPI.FetchPlanets(p.Next)
		if err != nil {
			panic(err)
		}
		planets = append(planets, p.Results...)
	}

	var films []*dtos.FilmsOutput

	for _, p := range planets {
		planet, _ := entities.NewPlanet(p.Name, p.Climate, p.Terrain)
		_, err = u.PlanetRepository.AddPlanet(planet)

		for _, film := range p.Films {
			m, err := u.StarWarsAPI.FetchFilms(film)
			if err != nil {
				panic(err)
			}
			f, _ := entities.NewFilm(planet.ID.String(), m.Title, m.Director, m.ReleaseDate)
			_ = u.PlanetRepository.AddFilm(f)

			films = append(films, m)
		}
		fmt.Printf("[Nome] [%s] [Clima] [%s] [Terreno] [%s]\n", p.Name, p.Climate, p.Terrain)
	}

	return nil
}

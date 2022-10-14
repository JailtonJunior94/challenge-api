package main

import (
	"fmt"

	"github.com/jailtonjunior94/challenge/internal/domain/dtos"
	"github.com/jailtonjunior94/challenge/internal/infra/facades"
)

func main() {
	starWarsAPI := facades.NewStarWarsFacade()

	p, err := starWarsAPI.FetchPlanets("https://swapi.dev/api/planets?page=1")
	if err != nil {
		panic(err)
	}

	var planets []dtos.PlanetsOutput
	planets = append(planets, p.Results...)

	for p.Next != "" {
		p, err = starWarsAPI.FetchPlanets(p.Next)
		if err != nil {
			panic(err)
		}
		planets = append(planets, p.Results...)
	}

	var films []*dtos.FilmsOutput

	for _, p := range planets {
		for _, film := range p.Films {
			m, err := starWarsAPI.FetchFilms(film)
			if err != nil {
				panic(err)
			}
			films = append(films, m)
		}
		fmt.Printf("[Nome] [%s] [Clima] [%s] [Terreno] [%s]\n", p.Name, p.Climate, p.Terrain)
	}
}

package usecases

import (
	"fmt"

	"github.com/jailtonjunior94/challenge/internal/domain/dtos"
	"github.com/jailtonjunior94/challenge/internal/domain/entities"
	"github.com/jailtonjunior94/challenge/internal/domain/interfaces"
	"github.com/jailtonjunior94/challenge/internal/infra/facades"

	"github.com/sirupsen/logrus"
)

type importPlanetUseCase struct {
	PlanetRepository interfaces.PlanetRepository
	StarWarsAPI      *facades.StarWarsFacade
	StarWarsBaseURL  string
}

func NewImportPlanetUseCase(planetRepository interfaces.PlanetRepository, starWarsAPI *facades.StarWarsFacade, starWarsBaseURL string) *importPlanetUseCase {
	return &importPlanetUseCase{
		PlanetRepository: planetRepository,
		StarWarsAPI:      starWarsAPI,
		StarWarsBaseURL:  starWarsBaseURL,
	}
}

func (u *importPlanetUseCase) Execute() error {
	planets, err := u.fetchPlanets()
	if err != nil {
		logrus.Errorf("Não foi possível obtér planetas, %v", err)
	}

	for _, p := range planets {
		err = u.PlanetRepository.AddPlanet(&p)
		if err != nil {
			logrus.Errorf("Erro ao inserir planeta %v", err)
			continue
		}

		for _, f := range p.Films {
			err = u.PlanetRepository.AddFilm(&f)
			if err != nil {
				logrus.Errorf("Erro ao inserir filme %v", err)
				continue
			}
		}
	}

	logrus.Info("Sucesso ao importar dados da API Star Wars")
	return nil
}

func (u *importPlanetUseCase) fetchPlanets() ([]entities.Planet, error) {
	p, err := u.StarWarsAPI.FetchPlanets(fmt.Sprintf("%s/planets?page=1", u.StarWarsBaseURL))
	if err != nil {
		return nil, err
	}

	var planets []dtos.PlanetsOutput
	planets = append(planets, p.Results...)

	for p.Next != "" {
		p, err = u.StarWarsAPI.FetchPlanets(p.Next)
		if err != nil {
			return nil, err
		}
		planets = append(planets, p.Results...)
	}

	var planetsEntity []entities.Planet

	for _, p := range planets {
		planet, err := entities.NewPlanet(p.Name, p.Climate, p.Terrain)
		if err != nil {
			logrus.Errorf("Erro ao gerar entidade de planetas %v", err)
			continue
		}

		films, err := u.fetchFilms(planet.ID, p.Films)
		if err != nil {
			logrus.Errorf("Erro ao obter filmes %v", err)
			continue
		}

		planet.AddFilms(films)
		planetsEntity = append(planetsEntity, *planet)
		logrus.Infof("Sucesso ao obter detalhes do planeta: %s", p.Name)
	}

	return planetsEntity, nil
}

func (u *importPlanetUseCase) fetchFilms(planetID string, filmsInput []string) ([]entities.Film, error) {
	var films []entities.Film

	for _, film := range filmsInput {
		f, err := u.StarWarsAPI.FetchFilms(film)
		if err != nil {
			logrus.Errorf("Erro ao obter detalhes do filme %v", err)
			continue
		}

		film, err := entities.NewFilm(planetID, f.Title, f.Director, f.ReleaseDate)
		if err != nil {
			logrus.Errorf("Erro ao gerar entidade de filmes %v", err)
			continue
		}

		films = append(films, *film)
		logrus.Infof("Sucesso ao obter detalhes do filme: %s", film.Title)
	}

	return films, nil
}

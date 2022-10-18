package usecases

import (
	"fmt"

	"github.com/jailtonjunior94/challenge/internal/domain/dtos"
	"github.com/jailtonjunior94/challenge/internal/domain/entities"
	"github.com/jailtonjunior94/challenge/internal/domain/interfaces"

	"github.com/sirupsen/logrus"
)

type importPlanetUseCase struct {
	PlanetRepository interfaces.PlanetRepository
	StarWarsAPI      interfaces.StarWarsFacade
	StarWarsBaseURL  string
}

func NewImportPlanetUseCase(planetRepository interfaces.PlanetRepository, starWarsAPI interfaces.StarWarsFacade, starWarsBaseURL string) *importPlanetUseCase {
	return &importPlanetUseCase{
		PlanetRepository: planetRepository,
		StarWarsAPI:      starWarsAPI,
		StarWarsBaseURL:  starWarsBaseURL,
	}
}

func (u *importPlanetUseCase) Execute() error {
	count, _ := u.PlanetRepository.CountPlanets("")
	if count > 0 {
		logrus.Info("[ImportPlanetUseCase] [Execute] [Base de dados já carregada]")
		return nil
	}

	planets, err := u.fetchPlanets()
	if err != nil {
		logrus.Errorf("[ImportPlanetUseCase] [Execute] [FetchPlanetsError] [Error] [%v]", err)
		return err
	}

	for _, p := range planets {
		err = u.PlanetRepository.AddPlanet(&p)
		if err != nil {
			logrus.Errorf("[ImportPlanetUseCase] [Execute] [AddPlanetError] [Error] [%v]", err)
			continue
		}

		for _, f := range p.Films {
			err = u.PlanetRepository.AddFilm(&f)
			if err != nil {
				logrus.Errorf("[ImportPlanetUseCase] [Execute] [AddFilmError] [Error] [%v]", err)
				continue
			}
		}
	}

	logrus.Info("[ImportPlanetUseCase] [Execute] [Sucesso ao importar dados da API Star Wars]")
	return nil
}

func (u *importPlanetUseCase) fetchPlanets() ([]entities.Planet, error) {
	p, err := u.StarWarsAPI.FetchPlanets(fmt.Sprintf("%s/planets?page=1", u.StarWarsBaseURL))
	if err != nil {
		logrus.Errorf("[ImportPlanetUseCase] [fetchPlanets] [Error] [%v]", err)
		return nil, err
	}

	var planets []dtos.PlanetsOutput
	planets = append(planets, p.Results...)

	for p.Next != "" {
		p, err = u.StarWarsAPI.FetchPlanets(p.Next)
		if err != nil {
			logrus.Errorf("[ImportPlanetUseCase] [fetchPlanets] [Error] [%v]", err)
			return nil, err
		}
		planets = append(planets, p.Results...)
	}

	var planetsEntity []entities.Planet

	for _, p := range planets {
		planet, err := entities.NewPlanet(p.Name, p.Climate, p.Terrain)
		if err != nil {
			logrus.Errorf("[ImportPlanetUseCase] [fetchPlanets] [NewPlanetError] [Error] [%v]", err)
			continue
		}

		films, _ := u.fetchFilms(planet.ID, p.Films)
		planet.AddFilms(films)
		planetsEntity = append(planetsEntity, *planet)
		logrus.Infof("[ImportPlanetUseCase] [fetchPlanets] [Sucesso ao obter detalhes do planeta] [%s]", p.Name)
	}

	return planetsEntity, nil
}

func (u *importPlanetUseCase) fetchFilms(planetID string, filmsInput []string) ([]entities.Film, error) {
	var films []entities.Film

	for _, film := range filmsInput {
		f, err := u.StarWarsAPI.FetchFilms(film)
		if err != nil {
			logrus.Errorf("[ImportPlanetUseCase] [fetchFilms] [Error] [%v]", err)
			continue
		}

		film, err := entities.NewFilm(planetID, f.Title, f.Director, f.ReleaseDate)
		if err != nil {
			logrus.Errorf("[ImportPlanetUseCase] [fetchFilms] [NewFilmError] [Error] [%v]", err)
			continue
		}

		films = append(films, *film)
		logrus.Infof("[ImportPlanetUseCase] [fetchFilms] [Sucesso ao obter detalhes do filme] [%s]", film.Title)
	}

	return films, nil
}

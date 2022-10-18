package facades

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/jailtonjunior94/challenge/internal/domain/dtos"

	"github.com/sirupsen/logrus"
)

var (
	ErrFetchPlanets = errors.New("cannot fetch planets")
	ErrFetchMovies  = errors.New("cannot fetch movies")
)

type StarWarsFacade struct {
	HttpClient *http.Client
}

func NewStarWarsFacade() *StarWarsFacade {
	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	return &StarWarsFacade{HttpClient: client}
}

func (f *StarWarsFacade) FetchPlanets(uri string) (*dtos.PaginateOutput[dtos.PlanetsOutput], error) {
	req, err := http.NewRequest(http.MethodGet, uri, nil)
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		logrus.Errorf("[StarWarsFacade] [FetchPlanets] [Error] [%v]", err)
		return nil, err
	}

	resp, err := f.HttpClient.Do(req)
	if err != nil {
		logrus.Errorf("[StarWarsFacade] [FetchPlanets] [Error] [%v]", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		logrus.Errorf("[StarWarsFacade] [FetchPlanets] [StatusCode] [%d]", resp.StatusCode)
		return nil, ErrFetchPlanets
	}

	var paginate *dtos.PaginateOutput[dtos.PlanetsOutput]
	if err := json.NewDecoder(resp.Body).Decode(&paginate); err != nil {
		return nil, err
	}

	logrus.Infof("[StarWarsFacade] [FetchPlanets] [Sucesso ao obter planeta] [Planeta] [%s]", paginate.Results[0].Name)
	return paginate, nil
}

func (f *StarWarsFacade) FetchFilms(uri string) (*dtos.FilmsOutput, error) {
	req, err := http.NewRequest(http.MethodGet, uri, nil)
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		logrus.Errorf("[StarWarsFacade] [FetchFilms] [Error] [%v]", err)
		return nil, err
	}

	resp, err := f.HttpClient.Do(req)
	if err != nil {
		logrus.Errorf("[StarWarsFacade] [FetchFilms] [Error] [%v]", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		logrus.Errorf("[StarWarsFacade] [FetchFilms] [StatusCode] [%d]", resp.StatusCode)
		return nil, ErrFetchMovies
	}

	var movie *dtos.FilmsOutput
	if err := json.NewDecoder(resp.Body).Decode(&movie); err != nil {
		return nil, err
	}

	logrus.Infof("[StarWarsFacade] [FetchFilms] [Sucesso ao obter filme] [Filme] [%s]", movie.Title)
	return movie, nil
}

package facades

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/jailtonjunior94/challenge/internal/domain/dtos"

	"github.com/sirupsen/logrus"
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
	resp, err := f.request(http.MethodGet, uri, "application/json")
	if err != nil {
		return nil, err
	}

	var paginate *dtos.PaginateOutput[dtos.PlanetsOutput]
	err = json.Unmarshal(resp, &paginate)
	if err != nil {
		return nil, err
	}

	return paginate, nil
}

func (f *StarWarsFacade) FetchFilms(uri string) (*dtos.FilmsOutput, error) {
	resp, err := f.request(http.MethodGet, uri, "application/json")
	if err != nil {
		return nil, err
	}

	var film *dtos.FilmsOutput
	err = json.Unmarshal(resp, &film)
	if err != nil {
		return nil, err
	}

	return film, nil
}

func (f *StarWarsFacade) request(method, uri, contentType string) ([]byte, error) {
	req, err := http.NewRequest(method, uri, nil)
	if err != nil {
		logrus.Errorf("[StarWarsFacade] [Error] [%v]", err)
		return nil, err
	}

	req.Header.Set("Content-Type", contentType)
	resp, err := f.HttpClient.Do(req)
	statusCode := resp.StatusCode
	if err != nil {
		logrus.Errorf("[StarWarsFacade] [Error] [%v]", err)
		return nil, err
	}

	if resp != nil {
		defer resp.Body.Close()
	}

	if statusCode < 200 || statusCode > 299 {
		b, _ := ioutil.ReadAll(resp.Body)
		return nil, errors.New(fmt.Sprintf("[ERROR] [StatusCode] [%d] [Detail] [%s]", statusCode, string(b)))
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	return bytes, err
}

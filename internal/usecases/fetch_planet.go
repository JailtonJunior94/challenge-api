package usecases

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/jailtonjunior94/challenge/internal/domain/dtos"
	"github.com/jailtonjunior94/challenge/internal/domain/interfaces"
	"github.com/jailtonjunior94/challenge/pkg/responses"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

type fetchPlanetUseCase struct {
	PlanetRepository interfaces.PlanetRepository
}

func NewFetchPlanetUseCase(planetRepository interfaces.PlanetRepository) *fetchPlanetUseCase {
	return &fetchPlanetUseCase{
		PlanetRepository: planetRepository,
	}
}

func (h *fetchPlanetUseCase) GetPlanetByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		responses.Error(w, http.StatusUnprocessableEntity, "ID ausente ou mal formatado")
		return
	}

	planet, err := h.PlanetRepository.FindByID(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			responses.Error(w, http.StatusNotFound, "Não foi encontrado nenhum planeta")
			logrus.Errorf("[FetchPlanetUseCase] [GetPlanetByID] [Error] [%v]", err)
			return
		}

		responses.Error(w, http.StatusInternalServerError, "Não foi possível encontrar planeta")
		logrus.Errorf("[FetchPlanetUseCase] [GetPlanetByID] [Error] [%v]", err)
		return
	}

	logrus.Info("[FetchPlanetUseCase] [GetPlanetByID] [Sucesso ao obter planeta por ID]")
	responses.JSON(w, http.StatusOK, planet)
}

func (h *fetchPlanetUseCase) GetPlanets(w http.ResponseWriter, r *http.Request) {
	input := dtos.NewFilterPlanetInput(r.URL.Query().Get("name"), r.URL.Query().Get("page"), r.URL.Query().Get("limit"))

	count, planets, err := h.PlanetRepository.FindAll(input.Name, input.Page, input.Limit)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			responses.Error(w, http.StatusNotFound, "Não foi encontrado nenhum planeta")
			logrus.Errorf("[FetchPlanetUseCase] [GetPlanets] [Error] [%v]", err)
			return
		}

		responses.Error(w, http.StatusInternalServerError, "Não foi possível encontrar planetas")
		logrus.Errorf("[FetchPlanetUseCase] [GetPlanets] [Error] [%v]", err)
		return
	}

	response := dtos.NewPaginationOutput(r.URL.Query().Get("page"), count, input.Limit, planets)
	logrus.Info("[FetchPlanetUseCase] [GetPlanets] [Sucesso ao obter planetas]")
	responses.JSON(w, http.StatusOK, response)
}

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

type FetchHandler struct {
	PlanetRepository interfaces.PlanetRepository
}

func NewFetchHandler(planetRepository interfaces.PlanetRepository) *FetchHandler {
	return &FetchHandler{
		PlanetRepository: planetRepository,
	}
}

func (h *FetchHandler) GetPlanetByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		responses.Error(w, http.StatusUnprocessableEntity, "ID ausente ou mal formatado")
		return
	}

	planet, err := h.PlanetRepository.FindByID(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			responses.Error(w, http.StatusNotFound, "Não foi encontrado nenhum planeta")
			logrus.Errorf(err.Error())
			return
		}

		responses.Error(w, http.StatusInternalServerError, "Não foi possível encontrar planeta")
		logrus.Errorf(err.Error())
		return
	}

	responses.JSON(w, http.StatusOK, planet)
}

func (h *FetchHandler) GetPlanets(w http.ResponseWriter, r *http.Request) {
	input := dtos.NewFilterPlanetInput(r.URL.Query().Get("name"), r.URL.Query().Get("page"), r.URL.Query().Get("limit"))

	planets, err := h.PlanetRepository.FindAll(input.Name, input.Page, input.Limit)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			responses.Error(w, http.StatusNotFound, "Não foi encontrado nenhum planeta")
			logrus.Errorf(err.Error())
			return
		}

		responses.Error(w, http.StatusInternalServerError, "Não foi possível encontrar planetas")
		logrus.Errorf(err.Error())
		return
	}

	responses.JSON(w, http.StatusOK, planets)
}

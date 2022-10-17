package usecases

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/jailtonjunior94/challenge/internal/domain/interfaces"
	"github.com/jailtonjunior94/challenge/pkg/responses"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

type RemoveHandler struct {
	PlanetRepository interfaces.PlanetRepository
}

func NewRemoveHandler(planetRepository interfaces.PlanetRepository) *RemoveHandler {
	return &RemoveHandler{
		PlanetRepository: planetRepository,
	}
}

func (h *RemoveHandler) RemovePlanetByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		responses.Error(w, http.StatusUnprocessableEntity, "ID ausente ou mal formatado")
		return
	}

	err := h.PlanetRepository.Remove(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			responses.Error(w, http.StatusNotFound, "Não foi encontrado nenhum planeta")
			logrus.Errorf(err.Error())
			return
		}

		responses.Error(w, http.StatusInternalServerError, "Não foi possível remover planeta")
		logrus.Errorf(err.Error())
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

package usecases

import (
	"encoding/json"
	"net/http"

	"github.com/jailtonjunior94/challenge/internal/domain/interfaces"

	"github.com/go-chi/chi/v5"
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
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	planet, err := h.PlanetRepository.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(planet)
}

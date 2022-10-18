package handlers

import (
	"net/http"

	"github.com/jailtonjunior94/challenge/internal/domain/dtos"
	"github.com/jailtonjunior94/challenge/internal/usecases"
	"github.com/jailtonjunior94/challenge/pkg/responses"

	"github.com/go-chi/chi/v5"
)

type PlanetHandler struct {
	RemoveUseCase   *usecases.RemovePlanetUseCase
	FindByIdUseCase *usecases.FindByIDUseCase
	FindAllUseCase  *usecases.FindAllUseCase
}

func NewPlanetHandler(removeUseCase *usecases.RemovePlanetUseCase,
	findByIDUseCase *usecases.FindByIDUseCase,
	findAllUseCase *usecases.FindAllUseCase) *PlanetHandler {
	return &PlanetHandler{
		RemoveUseCase:   removeUseCase,
		FindByIdUseCase: findByIDUseCase,
		FindAllUseCase:  findAllUseCase,
	}
}

func (h *PlanetHandler) GetPlanets(w http.ResponseWriter, r *http.Request) {
	input := dtos.NewFilterPlanetInput(r.URL.Query().Get("name"), r.URL.Query().Get("page"), r.URL.Query().Get("limit"))
	response := h.FindAllUseCase.Execute(input)
	responses.JSON(w, response.StatusCode, response.Data)
}

func (h *PlanetHandler) GetPlanetByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		responses.Error(w, http.StatusUnprocessableEntity, "ID ausente ou mal formatado")
		return
	}

	response := h.FindByIdUseCase.Execute(id)
	responses.JSON(w, response.StatusCode, response.Data)
}

func (h *PlanetHandler) RemovePlanet(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		responses.Error(w, http.StatusUnprocessableEntity, "ID ausente ou mal formatado")
		return
	}

	response := h.RemoveUseCase.Execute(id)
	responses.JSON(w, response.StatusCode, response.Data)
}

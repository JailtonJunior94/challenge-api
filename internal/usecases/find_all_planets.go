package usecases

import (
	"net/http"

	"github.com/jailtonjunior94/challenge/internal/domain/dtos"
	"github.com/jailtonjunior94/challenge/internal/domain/interfaces"
	"github.com/sirupsen/logrus"
)

type FindAllUseCase struct {
	PlanetRepository interfaces.PlanetRepository
}

func NewFindAllUseCase(planetRepository interfaces.PlanetRepository) *FindAllUseCase {
	return &FindAllUseCase{
		PlanetRepository: planetRepository,
	}
}

func (h *FindAllUseCase) Execute(input *dtos.FilterPlanetInput) *dtos.HttpResponse {
	count, planets, err := h.PlanetRepository.FindAll(input)
	if count <= 0 {
		logrus.Errorf("[FindAllUseCase] [Error] [%v]", err)
		return dtos.NewHttpResponse(http.StatusNoContent, nil)
	}

	if err != nil {
		logrus.Errorf("[FindAllUseCase] [Error] [%v]", err)
		return dtos.NewHttpResponse(http.StatusInternalServerError, nil)
	}

	response := dtos.NewPaginationOutput(count, input.Limit, planets)
	logrus.Info("[FindAllUseCase] [Planetas encontrados com sucesso]")
	return dtos.NewHttpResponse(http.StatusOK, response)
}

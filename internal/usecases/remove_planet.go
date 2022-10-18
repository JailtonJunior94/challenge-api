package usecases

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/jailtonjunior94/challenge/internal/domain/dtos"
	"github.com/jailtonjunior94/challenge/internal/domain/interfaces"

	"github.com/sirupsen/logrus"
)

type RemovePlanetUseCase struct {
	PlanetRepository interfaces.PlanetRepository
}

func NewRemovePlanetUseCase(planetRepository interfaces.PlanetRepository) *RemovePlanetUseCase {
	return &RemovePlanetUseCase{
		PlanetRepository: planetRepository,
	}
}

func (u *RemovePlanetUseCase) Execute(id string) *dtos.HttpResponse {
	err := u.PlanetRepository.Remove(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			logrus.Errorf("[RemovePlanetUseCase] [Error] [%v]", err)
			return dtos.NewHttpResponse(http.StatusNoContent, nil)
		}
		logrus.Errorf("[RemovePlanetUseCase] [Error] [%v]", err)
		return dtos.NewHttpResponse(http.StatusInternalServerError, nil)
	}

	logrus.Info("[RemovePlanetUseCase] [Planeta removido com sucesso]")
	return dtos.NewHttpResponse(http.StatusNoContent, nil)
}

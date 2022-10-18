package usecases

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/jailtonjunior94/challenge/internal/domain/dtos"
	"github.com/jailtonjunior94/challenge/internal/domain/interfaces"

	"github.com/sirupsen/logrus"
)

type FindByIDUseCase struct {
	PlanetRepository interfaces.PlanetRepository
}

func NewFindByIDUseCase(planetRepository interfaces.PlanetRepository) *FindByIDUseCase {
	return &FindByIDUseCase{
		PlanetRepository: planetRepository,
	}
}

func (u *FindByIDUseCase) Execute(id string) *dtos.HttpResponse {
	planet, err := u.PlanetRepository.FindByID(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			logrus.Errorf("[FindByIDUseCase] [Error] [%v]", err)
			return dtos.NewHttpResponse(http.StatusNoContent, nil)
		}
		logrus.Errorf("[FindByIDUseCase] [Error] [%v]", err)
		return dtos.NewHttpResponse(http.StatusInternalServerError, nil)
	}

	logrus.Info("[FindByIDUseCase] [Planeta encontrado com sucesso]")
	return dtos.NewHttpResponse(http.StatusOK, planet)
}

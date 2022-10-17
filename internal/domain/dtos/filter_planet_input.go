package dtos

import "strconv"

type FilterPlanetInput struct {
	Name  string `json:"name"`
	Page  int    `json:"page"`
	Limit int    `json:"limit"`
}

func NewFilterPlanetInput(name, page, limit string) *FilterPlanetInput {
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		pageInt = 0
	}

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		limitInt = 0
	}

	return &FilterPlanetInput{
		Name:  name,
		Page:  (pageInt - 1) * limitInt,
		Limit: limitInt,
	}
}

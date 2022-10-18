package dtos

import (
	"math"
	"strconv"
)

type PaginationOutput[T any] struct {
	Page       int `json:"page"`
	TotalPages int `json:"totalPages"`
	TotalItems int `json:"totalItems"`
	Items      []T `json:"items"`
}

func NewPaginationOutput[T any](page string, count, limit int, items []T) *PaginationOutput[T] {
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		pageInt = 0
	}

	return &PaginationOutput[T]{
		Page:       pageInt,
		TotalPages: int(math.Ceil(float64(count) / float64(limit))),
		TotalItems: count,
		Items:      items,
	}
}

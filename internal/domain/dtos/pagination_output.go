package dtos

import (
	"math"
)

type PaginationOutput[T any] struct {
	TotalPages int `json:"totalPages"`
	TotalItems int `json:"totalItems"`
	Items      []T `json:"items"`
}

func NewPaginationOutput[T any](count, limit int, items []T) *PaginationOutput[T] {
	return &PaginationOutput[T]{

		TotalPages: int(math.Ceil(float64(count) / float64(limit))),
		TotalItems: count,
		Items:      items,
	}
}

package dtos

type PaginateOutput[T any] struct {
	Count    int         `json:"count"`
	Next     string      `json:"next"`
	Previous interface{} `json:"previous"`
	Results  []T         `json:"results"`
}

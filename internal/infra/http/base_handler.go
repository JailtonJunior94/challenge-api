package http

import "database/sql"

type baseHandler struct {
	DB *sql.DB
}

func NewBaseHandler(db *sql.DB) *baseHandler {
	return &baseHandler{DB: db}
}

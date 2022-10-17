package entity

import uuid "github.com/satori/go.uuid"

type ID = uuid.UUID

func NewID() ID {
	return ID(uuid.NewV4())
}

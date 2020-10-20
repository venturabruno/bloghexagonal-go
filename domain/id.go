package domain

import "github.com/google/uuid"

type EntityID = uuid.UUID

func NewID() EntityID {
	return EntityID(uuid.New())
}

func StringToEntityID(s string) (EntityID, error) {
	id, err := uuid.Parse(s)
	return EntityID(id), err
}

package data

import "github.com/gofrs/uuid"

type EntityReader[ENTITY interface{}] interface {
	ReadOne(id uuid.UUID) (*ENTITY, error)
	ReadAll() ([]*ENTITY, error)
}

type EntityWriter[ENTITY interface{}] interface {
	CreateOne(entity *ENTITY) (*ENTITY, error)
}

type EntityAccess[ENTITY interface{}] struct {
	Reader EntityReader[ENTITY]
	Writer EntityWriter[ENTITY]
}

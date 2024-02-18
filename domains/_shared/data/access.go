package data

import "github.com/gofrs/uuid"

type PageInfo struct {
	PageSize   int32
	PageOffset int32
}

type EntityReader[ENTITY interface{}] interface {
	ReadOne(id uuid.UUID) (*ENTITY, error)
	ReadMany(pageInfo PageInfo) ([]*ENTITY, error)
}

type EntityWriter[ENTITY interface{}] interface {
	CreateOne(entity *ENTITY) (*ENTITY, error)
}

type EntityAccess[ENTITY interface{}] struct {
	Reader EntityReader[ENTITY]
	Writer EntityWriter[ENTITY]
}

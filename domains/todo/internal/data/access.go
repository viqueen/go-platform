package data

import "github.com/gofrs/uuid"

type Todo struct {
	Id    uuid.UUID
	Title string
}

type TodoReader interface {
	ReadOne(id uuid.UUID) (*Todo, error)
}

type TodoWriter interface {
	CreateOne(todo Todo) (*Todo, error)
}

type TodoAccess struct {
	Reader TodoReader
	Writer TodoWriter
}

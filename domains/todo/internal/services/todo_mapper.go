package services

import (
	todoV1 "github.com/viqueen/go-platform/domains/_api/todo/v1"
	"github.com/viqueen/go-platform/domains/todo/internal/data"
)

func modelToApi(model *data.Todo) *todoV1.Todo {
	return &todoV1.Todo{
		Id:    model.Id.String(),
		Title: model.Title,
	}
}

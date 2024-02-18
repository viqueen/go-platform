package data

import (
	todoV1 "github.com/viqueen/go-platform/domains/_api/todo/v1"
	"github.com/viqueen/go-platform/domains/_shared/data"
)

type TodoAccess struct {
	Todo *data.EntityAccess[todoV1.Todo]
}

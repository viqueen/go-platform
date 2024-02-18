package neo4j_access

import (
	"github.com/neo4j/neo4j-go-driver/neo4j"
	todoV1 "github.com/viqueen/go-platform/domains/_api/todo/v1"
	"github.com/viqueen/go-platform/domains/_shared/clients"
	sharedData "github.com/viqueen/go-platform/domains/_shared/data"
	sharedNeo4jAccess "github.com/viqueen/go-platform/domains/_shared/data/neo4j-access"
	"github.com/viqueen/go-platform/domains/todo/internal/data"
)

func NewTodoAccess(neo4jClient *clients.Neo4jClient) *data.TodoAccess {
	return &data.TodoAccess{
		Todo: &sharedData.EntityAccess[todoV1.Todo]{
			Reader: sharedNeo4jAccess.NewEntityNeo4jReader[todoV1.Todo](
				neo4jClient,
				"Todo",
				[]string{"id", "title"},
				func(record neo4j.Record) *todoV1.Todo {
					return &todoV1.Todo{
						Id:    record.GetByIndex(0).(string),
						Title: record.GetByIndex(1).(string),
					}
				},
			),
			Writer: sharedNeo4jAccess.NewEntityNeo4jWriter[todoV1.Todo](
				neo4jClient,
				"Todo",
				[]string{"id", "title"},
				func(entity *todoV1.Todo) map[string]interface{} {
					return map[string]interface{}{
						"id":    entity.GetId(),
						"title": entity.GetTitle(),
					}
				},
			),
		},
	}
}

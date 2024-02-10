package neo4j_access

import (
	"github.com/viqueen/go-platform/domains/_shared/clients"
	"github.com/viqueen/go-platform/domains/todo/internal/data"
)

type TodoNeo4jWriter struct {
	data.TodoWriter
	client *clients.Neo4jClient
}

func (w TodoNeo4jWriter) CreateOne(todo data.Todo) (*data.Todo, error) {
	query := "CREATE (t:Todo {id: $id, title: $title})"
	params := map[string]interface{}{
		"id":    todo.Id.String(),
		"title": todo.Title,
	}

	err := w.client.ExecuteWriteQuery(query, params)
	if err != nil {
		return nil, err
	}

	return &todo, nil
}

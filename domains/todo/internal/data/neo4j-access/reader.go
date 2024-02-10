package neo4j_access

import (
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/viqueen/go-platform/domains/_shared/clients"
	"github.com/viqueen/go-platform/domains/todo/internal/data"
)

type TodoNeo4jReader struct {
	data.TodoReader
	client *clients.Neo4jClient
}

func (r TodoNeo4jReader) ReadOne(id uuid.UUID) (*data.Todo, error) {
	query := "MATCH (t:Todo {id: $id}) RETURN t.id, t.title LIMIT 1"
	params := map[string]interface{}{
		"id": id.String(),
	}

	result, err := r.client.ExecuteReadQuery(query, params)
	if err != nil {
		return nil, err
	}

	hasRecord := result.Next()
	if !hasRecord {
		return nil, fmt.Errorf("content with ID %s not found", id)
	}

	record := result.Record()
	content := &data.Todo{
		Id:    uuid.FromStringOrNil(record.GetByIndex(0).(string)),
		Title: record.GetByIndex(1).(string),
	}

	return content, nil
}

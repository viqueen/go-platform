package neo4j_access

import (
	"github.com/viqueen/go-platform/domains/_shared/clients"
	"github.com/viqueen/go-platform/domains/todo/internal/data"
)

func NewTodoNeo4jAccess() (*data.TodoAccess, error) {
	neo4jClient, err := clients.NewLocalNeo4jClient()
	if err != nil {
		return nil, err
	}
	return &data.TodoAccess{
		Reader: &TodoNeo4jReader{client: neo4jClient},
		Writer: &TodoNeo4jWriter{client: neo4jClient},
	}, nil
}

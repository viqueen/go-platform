package export

import (
	todoV1 "github.com/viqueen/go-platform/domains/_api/todo/v1"
	neo4jaccess "github.com/viqueen/go-platform/domains/todo/internal/data/neo4j-access"
	"github.com/viqueen/go-platform/domains/todo/internal/services"
	"google.golang.org/grpc"
)

func Bundle(server *grpc.Server) error {
	todoAccess, err := neo4jaccess.NewTodoNeo4jAccess()
	if err != nil {
		return err
	}
	todoService, err := services.NewTodoService(todoAccess)
	todoV1.RegisterTodoServiceServer(server, todoService)
	return nil
}

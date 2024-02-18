package export

import (
	todoV1 "github.com/viqueen/go-platform/domains/_api/todo/v1"
	"github.com/viqueen/go-platform/domains/_shared/clients"
	neo4jaccess "github.com/viqueen/go-platform/domains/todo/internal/data/neo4j-access"
	"github.com/viqueen/go-platform/domains/todo/internal/services"
	"google.golang.org/grpc"
)

func Bundle(server *grpc.Server, neo4jClient *clients.Neo4jClient) {
	todoAccess := neo4jaccess.NewTodoAccess(neo4jClient)
	todoService := services.NewTodoService(todoAccess)
	todoV1.RegisterTodoServiceServer(server, todoService)
}

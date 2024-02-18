package dynamodb_access

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	todoV1 "github.com/viqueen/go-platform/domains/_api/todo/v1"
	sharedData "github.com/viqueen/go-platform/domains/_shared/data"
	sharedDynamoDbAccess "github.com/viqueen/go-platform/domains/_shared/data/dynamodb-access"
	"github.com/viqueen/go-platform/domains/todo/internal/data"
)

func NewTodoAccess(dynamoDbClient *dynamodb.DynamoDB) *data.TodoAccess {
	return &data.TodoAccess{
		Todo: &sharedData.EntityAccess[todoV1.Todo]{
			Reader: sharedDynamoDbAccess.NewEntityDynamoDBReader[todoV1.Todo](
				dynamoDbClient,
				"Todo",
			),
			Writer: sharedDynamoDbAccess.NewEntityDynamoDBWriter[todoV1.Todo](
				dynamoDbClient,
				"Todo",
			),
		},
	}
}

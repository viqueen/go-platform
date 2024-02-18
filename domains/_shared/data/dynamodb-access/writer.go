package dynamodb_access

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/viqueen/go-platform/domains/_shared/data"
)

type EntityDynamoDBWriter[ENTITY interface{}] struct {
	data.EntityWriter[ENTITY]

	client     *dynamodb.DynamoDB
	entityName string
}

func NewEntityDynamoDBWriter[ENTITY interface{}](
	client *dynamodb.DynamoDB,
	entityName string,
) *EntityDynamoDBWriter[ENTITY] {
	return &EntityDynamoDBWriter[ENTITY]{
		client:     client,
		entityName: entityName,
	}
}

func (w *EntityDynamoDBWriter[ENTITY]) CreateOne(entity *ENTITY) (*ENTITY, error) {
	item, err := dynamodbattribute.MarshalMap(entity)
	if err != nil {
		return nil, err
	}

	input := &dynamodb.PutItemInput{
		TableName: &w.entityName,
		Item:      item,
	}

	_, err = w.client.PutItem(input)
	if err != nil {
		return nil, err
	}

	return entity, nil
}

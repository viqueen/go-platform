package dynamodb_access

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/gofrs/uuid"
	"github.com/viqueen/go-platform/domains/_shared/data"
)

type EntityDynamoDBReader[ENTITY interface{}] struct {
	data.EntityReader[ENTITY]

	client     *dynamodb.DynamoDB
	entityName string
}

func NewEntityDynamoDBReader[ENTITY interface{}](client *dynamodb.DynamoDB, entityName string) *EntityDynamoDBReader[ENTITY] {
	return &EntityDynamoDBReader[ENTITY]{
		client:     client,
		entityName: entityName,
	}
}

func (r *EntityDynamoDBReader[ENTITY]) ReadOne(id uuid.UUID) (*ENTITY, error) {
	input := &dynamodb.GetItemInput{
		TableName: &r.entityName,
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id.String()),
			},
		},
	}

	result, err := r.client.GetItem(input)
	if err != nil {
		return nil, err
	}

	if len(result.Item) == 0 {
		return nil, fmt.Errorf("entity %s with ID %s not found", r.entityName, id)
	}

	var entity ENTITY
	err = dynamodbattribute.UnmarshalMap(result.Item, &entity)
	if err != nil {
		return nil, err
	}

	return &entity, nil
}

func (r *EntityDynamoDBReader[ENTITY]) ReadMany(_ data.PageInfo) ([]*ENTITY, error) {
	panic("implement	me")
}

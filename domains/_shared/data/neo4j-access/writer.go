package neo4j_access

import (
	"fmt"
	"github.com/viqueen/go-platform/domains/_shared/clients"
	"github.com/viqueen/go-platform/domains/_shared/collections"
	"github.com/viqueen/go-platform/domains/_shared/data"
	"strings"
)

type EntityNeo4jWriter[ENTITY interface{}] struct {
	data.EntityWriter[ENTITY]
	client *clients.Neo4jClient

	entityName   string
	entityFields []string
	paramsMapper func(entity *ENTITY) map[string]interface{}
}

func NewEntityNeo4jWriter[ENTITY interface{}](
	client *clients.Neo4jClient,
	entityName string,
	entityFields []string,
	paramsMapper func(entity *ENTITY) map[string]interface{},
) *EntityNeo4jWriter[ENTITY] {
	return &EntityNeo4jWriter[ENTITY]{
		client:       client,
		entityName:   entityName,
		entityFields: entityFields,
		paramsMapper: paramsMapper,
	}
}

func (w *EntityNeo4jWriter[ENTITY]) CreateOne(entity *ENTITY) (*ENTITY, error) {
	fields := collections.Map(w.entityFields, func(field string) string {
		return fmt.Sprintf("%s: $%s", field, field)
	})
	joined := strings.Join(fields, ", ")
	query := fmt.Sprintf("CREATE (t:%s {%s})", w.entityName, joined)
	params := w.paramsMapper(entity)
	err := w.client.ExecuteWriteQuery(query, params)

	if err != nil {
		return nil, err
	}

	return entity, nil
}

package neo4j_access

import (
	"github.com/viqueen/go-music/domains/_shared/clients"
	"github.com/viqueen/go-music/domains/_shared/data"
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
	params := w.paramsMapper(entity)
	query := "CREATE (t:%s) SET %s RETURN t"
	err := w.client.ExecuteWriteQuery(query, params)

	if err != nil {
		return nil, err
	}

	return entity, nil
}

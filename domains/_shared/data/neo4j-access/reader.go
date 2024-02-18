package neo4j_access

import (
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/neo4j/neo4j-go-driver/neo4j"
	"github.com/viqueen/go-platform/domains/_shared/clients"
	"github.com/viqueen/go-platform/domains/_shared/collections"
	"github.com/viqueen/go-platform/domains/_shared/data"
	"strings"
)

type EntityNeo4jReader[ENTITY interface{}] struct {
	data.EntityReader[ENTITY]

	client       *clients.Neo4jClient
	entityName   string
	entityFields []string
	recordMapper func(record neo4j.Record) *ENTITY
}

func NewEntityNeo4jReader[ENTITY interface{}](
	client *clients.Neo4jClient,
	entityName string,
	entityFields []string,
	recordMapper func(record neo4j.Record) *ENTITY,
) *EntityNeo4jReader[ENTITY] {
	return &EntityNeo4jReader[ENTITY]{
		client:       client,
		entityName:   entityName,
		entityFields: entityFields,
		recordMapper: recordMapper,
	}
}

func (r *EntityNeo4jReader[ENTITY]) ReadOne(id uuid.UUID) (*ENTITY, error) {
	fields := collections.Map(r.entityFields, func(field string) string {
		return fmt.Sprintf("t.%s", field)
	})
	joined := strings.Join(fields, ", ")
	query := fmt.Sprintf("MATCH (t:%s {id: $id}) RETURN %s LIMIT 1", r.entityName, joined)
	params := map[string]interface{}{
		"id": id.String(),
	}

	result, err := r.client.ExecuteReadQuery(query, params)
	if err != nil {
		return nil, err
	}

	hasRecord := result.Next()
	if !hasRecord {
		return nil, fmt.Errorf("entity %s with ID %s not found", r.entityName, id)
	}

	record := result.Record()
	entity := r.recordMapper(record)

	return entity, nil
}

func (r *EntityNeo4jReader[ENTITY]) ReadMany(pageInfo data.PageInfo) ([]*ENTITY, error) {
	fields := collections.Map(r.entityFields, func(field string) string {
		return fmt.Sprintf("t.%s", field)
	})
	joined := strings.Join(fields, ", ")
	query := fmt.Sprintf("MATCH (t:%s) RETURN %s SKIP %d LIMIT %d",
		r.entityName,
		joined,
		pageInfo.PageOffset,
		pageInfo.PageSize,
	)
	result, err := r.client.ExecuteReadQuery(query, nil)
	if err != nil {
		return nil, err
	}

	var entities []*ENTITY
	for result.Next() {
		record := result.Record()
		entity := r.recordMapper(record)
		entities = append(entities, entity)
	}

	return entities, nil
}

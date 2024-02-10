package clients

import (
	"github.com/neo4j/neo4j-go-driver/neo4j"
	"log"
)

// Neo4jReader defines the interface for reading operations from Neo4j.
type Neo4jReader interface {
	ExecuteReadQuery(query string, params map[string]interface{}) (neo4j.Result, error)
}

// Neo4jWriter defines the interface for writing operations to Neo4j.
type Neo4jWriter interface {
	ExecuteWriteQuery(query string, params map[string]interface{}) (neo4j.Result, error)
}

type Neo4jClient struct {
	driver neo4j.Driver
}

func NewNeo4jClient(uri string, username string, password string) (*Neo4jClient, error) {
	driver, err := neo4j.NewDriver(uri, neo4j.BasicAuth(username, password, ""))
	if err != nil {
		return nil, err
	}
	return &Neo4jClient{
		driver: driver,
	}, nil
}

func NewLocalNeo4jClient() (*Neo4jClient, error) {
	neo4jURI := "bolt://localhost:7687"
	neo4jUsername := "neo4j"
	neo4jPassword := "password"
	driver, err := neo4j.NewDriver(neo4jURI, neo4j.BasicAuth(neo4jUsername, neo4jPassword, ""), func(c *neo4j.Config) {
		c.Encrypted = false
	})
	if err != nil {
		return nil, err
	}
	return &Neo4jClient{
		driver: driver,
	}, nil
}

func (n *Neo4jClient) ExecuteReadQuery(query string, params map[string]interface{}) (neo4j.Result, error) {
	session, _ := n.driver.NewSession(neo4j.SessionConfig{
		AccessMode:   neo4j.AccessModeRead,
		DatabaseName: "",
	})
	defer session.Close()

	result, err := session.ReadTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		return transaction.Run(query, params)
	})

	if err != nil {
		return nil, err
	}

	return result.(neo4j.Result), nil
}

func (n *Neo4jClient) ExecuteWriteQuery(query string, params map[string]interface{}) error {
	session, _ := n.driver.NewSession(neo4j.SessionConfig{
		AccessMode:   neo4j.AccessModeWrite,
		DatabaseName: "",
	})
	defer func(session neo4j.Session) {
		err := session.Close()
		if err != nil {
			log.Panicf("neo4j: failed session close - %v", err)
		}
	}(session)

	_, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		return transaction.Run(query, params)
	})

	if err != nil {
		return err
	}

	return nil
}

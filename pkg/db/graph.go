package db

import (
	"log"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"

	"github.com/SaidakbarPardaboyev/Mini-Tweeter-Users-Service/config"
	// "github.com/dgraph-io/dgo/v210"
	// "github.com/dgraph-io/dgo/v210/protos/api"
)

func NewDgraphClient(cfg *config.Config) (neo4j.Driver, neo4j.Session, error) {
	// Connect to Neo4j
	driver, err := neo4j.NewDriver("neo4j+s://f19fa800.databases.neo4j.io", neo4j.BasicAuth("neo4j", "_fo0c1CNc6A-4o8iA9cBPf4qmW22Qm4XEA2CepcS73M", ""))
	if err != nil {
		log.Fatal("Error connecting to Neo4j:", err)
	}

	// Create a session
	session := driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})

	return driver, session, nil
}

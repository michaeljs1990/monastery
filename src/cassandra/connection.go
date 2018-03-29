package cassandra

import (
	"github.com/gocql/gocql"

	"github.com/michaeljs1990/monastery/src/config"
)

// CreateSession creates a session to query cassandra through
func CreateSession() *gocql.Session {
	cluster := gocql.NewCluster(config.CQLServer)
	cluster.Keyspace = config.CQLKeyspace

	session, err := cluster.CreateSession()
	if err != nil {
		panic(err)
	}

	return session
}

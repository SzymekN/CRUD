package storage

import (
	"log"

	"github.com/gocql/gocql"
)

var CASSANDRA *gocql.Session

func CreateCassandraSession() *gocql.Session {

	var err error
	cluster := gocql.NewCluster("192.168.33.50")
	cluster.Keyspace = "userapi"
	cluster.Consistency = gocql.Quorum
	CASSANDRA, err = cluster.CreateSession()

	if err != nil {
		log.Fatal(err)
	}

	return CASSANDRA
}

func GetCassandraInstance() *gocql.Session {
	return CASSANDRA
}

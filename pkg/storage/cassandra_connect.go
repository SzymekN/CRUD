package storage

import (
	"log"
	"os"

	"github.com/gocql/gocql"
)

var CASSANDRA *gocql.Session

func SetupCassandraConnection() *gocql.Session {

	var err error

	cassIP := os.Getenv("CASSANDRA_HOST")

	// if cassIP == "" {
	// 	cassIP = "192.168.33.50"
	// }

	cluster := gocql.NewCluster(cassIP)
	// cluster.Keyspace = "userapi"
	cluster.Consistency = gocql.Quorum
	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: os.Getenv("CASSANDRA_USER"),
		Password: os.Getenv("CASSANDRA_PASSWORD"),
	}
	CASSANDRA, err = cluster.CreateSession()

	if err != nil {
		log.Fatal(err)
	}

	return CASSANDRA
}

func GetCassandraInstance() *gocql.Session {
	return CASSANDRA
}

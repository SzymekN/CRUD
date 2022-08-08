package storage

import (
	"log"

	config "crud/config"

	"github.com/gocql/gocql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB
var CASSANDRA *gocql.Session

func NewDB(params ...string) *gorm.DB {
	var err error
	conString := config.GetPostgresConnectionString()

	log.Print(conString)

	DB, err = gorm.Open(config.GetDBType(), conString)

	if err != nil {
		log.Panic(err)
	}

	return DB
}

func GetDBInstance() *gorm.DB {
	return DB
}

func CreateCassandraSession() *gocql.Session {

	var err error
	cluster := gocql.NewCluster("192.168.33.40")
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

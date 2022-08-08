package main

import (
	controller "crud/controller"
	"crud/storage"
	"fmt"
	"log"

	"github.com/gocql/gocql"
)

func connectionTest() {
	cluster := gocql.NewCluster("192.168.33.40")
	cluster.Keyspace = "store"
	cluster.Consistency = gocql.Quorum
	session, err := cluster.CreateSession()
	if err != nil {
		fmt.Println(err)
	}
	defer session.Close()

	if err := session.Query(`INSERT INTO store.shopping_cart (userid, item_count) VALUES ('4567', 20);`).Exec(); err != nil {
		log.Fatal(err)
	}

	var id string
	var count int

	if err := session.Query(` SELECT userid,item_count FROM store.shopping_cart limit 1;`).Consistency(gocql.One).Scan(&id, &count); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Tweet:", id, count)

	// iter := session.Query(`SELECT id, text FROM tweet WHERE timeline = ?`, "me").Iter()
	// for iter.Scan(&id, &text) {
	// 	fmt.Println("Tweet:", id, text)
	// }
	// if err := iter.Close(); err != nil {
	// 	log.Fatal(err)
	// }
}

func main() {

	e := controller.SetupRouter()
	// storage.NewDB()
	storage.CreateCassandraSession()
	port := ":8200"
	// if len(os.Args) > 1 {
	// 	port = os.Args[1]
	// } else {
	// 	if temp := os.Getenv("ECHO_PORT"); temp != "" {
	// 		port = temp
	// 	}
	// }

	// seeder.CreateAndSeed(storage.GetDBInstance())
	// connectionTest()

	e.Logger.Fatal(e.Start(port))
}

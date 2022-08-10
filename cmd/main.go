package main

import (
	"crud/pkg/controller"
	"crud/pkg/seeder"
	"crud/pkg/storage"
	"os"
)

func main() {

	e := controller.SetupRouter()
	storage.CreateCassandraSession()
	storage.NewDB()
	port := ":8200"
	if len(os.Args) > 1 {
		port = os.Args[1]
	} else {
		if temp := os.Getenv("ECHO_PORT"); temp != "" {
			port = temp
		}
	}

	// if len(os.Args) > 1 {
	seeder.CreateAndSeed(storage.GetDBInstance(), true)
	// }

	e.Logger.Fatal(e.Start(port))
}

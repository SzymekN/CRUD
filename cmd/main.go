package main

import (
	"crud/pkg/controller"
	"crud/pkg/seeder"
	"crud/pkg/storage"
)

func main() {

	e := controller.SetupRouter()
	storage.CreateCassandraSession()
	storage.NewDB()

	seeder.CreateAndSeed(storage.GetDBInstance(), true)

	e.Logger.Fatal(e.Start(":8200"))
}

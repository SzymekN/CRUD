package main

import (
	"crud/pkg/grpc"
	"crud/pkg/storage"
)

func main() {

	// e := controller.SetupRouter()
	storage.CreateCassandraSession()
	grpc.CreateGRPCServer()
	// storage.NewDB()

	// seeder.CreateAndSeed(storage.GetDBInstance(), true)

	// e.Logger.Fatal(e.Start(":8200"))

}

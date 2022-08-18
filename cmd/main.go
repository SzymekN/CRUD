package main

import (
	"github.com/SzymekN/CRUD/pkg/controller"
	"github.com/SzymekN/CRUD/pkg/grpc"
	"github.com/SzymekN/CRUD/pkg/storage"
)

func main() {

	e := controller.SetupRouter()
	storage.CreateCassandraSession()
	go grpc.CreateGRPCServer()
	storage.NewDB()

	// seeder.CreateAndSeed(storage.GetDBInstance(), true)

	e.Logger.Fatal(e.Start(":8200"))

}

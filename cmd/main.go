package main

import (
	"github.com/SzymekN/CRUD/pkg/controller"
	"github.com/SzymekN/CRUD/pkg/grpc"
	"github.com/SzymekN/CRUD/pkg/producer"
	"github.com/SzymekN/CRUD/pkg/storage"
)

func main() {

	e := controller.SetupRouter()
	storage.CreateCassandraSession()
	storage.NewDB()
	producer.Setup()

	// seeder.CreateAndSeed(storage.GetDBInstance(), true)

	go grpc.CreateGRPCServer()
	e.Logger.Fatal(e.Start(":8200"))

}

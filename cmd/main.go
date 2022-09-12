package main

import (
	"fmt"
	"os"

	"github.com/SzymekN/CRUD/pkg/controller"
	"github.com/SzymekN/CRUD/pkg/grpc"
	"github.com/SzymekN/CRUD/pkg/producer"
	"github.com/SzymekN/CRUD/pkg/seeder"
	"github.com/SzymekN/CRUD/pkg/storage"
)

func main() {

	e := controller.SetupRouter()
	storage.SetupRedisConnection()
	storage.SetupCassandraConnection()
	storage.SetupPostgresConnection()
	fmt.Println("ŁOTR 1222")
	producer.Setup()
	seeder.CreateAndSeed()
	// defer CloseAll(){}
	go grpc.CreateGRPCServer()
	e.Logger.Fatal(e.Start(":" + os.Getenv("API_PORT")))

}

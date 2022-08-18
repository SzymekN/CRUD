package main

import (
	// _ "github.com/SzymekN/CRUD/docs/userapi"
	"github.com/SzymekN/CRUD/pkg/controller"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func main() {

	e := controller.SetupRouter()
	// storage.CreateCassandraSession()
	// go grpc.CreateGRPCServer()
	// storage.NewDB()

	// seeder.CreateAndSeed(storage.GetDBInstance(), true)

	// opts := middleware.SwaggerUIOpts{SpecURL: "/swagger.yaml"}
	// sh := middleware.SwaggerUI(opts, nil)
	// e.Handle("/docs", sh)

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(":8200"))

}

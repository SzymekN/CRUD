package main

import (
	controller "crud/controller"
	storage "crud/storage"
	"os"
)

func main() {

	e := controller.SetupRouter()
	storage.NewDB()
	port := ":8200"
	if len(os.Args) > 1 {
		port = os.Args[1]
	} else {
		port = os.Getenv("ECHO_PORT")
	}
	e.Logger.Fatal(e.Start(port))
}

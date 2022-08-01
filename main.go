package main

import (
	controller "crud/controller"
	storage "crud/storage"
)

func main() {

	e := controller.SetupRouter()
	storage.NewDB()
	e.Logger.Fatal(e.Start(":8200")) // port as env variable
}

package main

import (
	storage "crud/database"
	server "crud/server"
)

func main() {

	e := server.SetupRouter()
	storage.NewDB()
	e.Logger.Fatal(e.Start(":8200")) // port as env variable
}

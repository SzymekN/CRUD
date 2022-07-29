package main

import (
	server "crud/server"
)

func main() {

	e := server.SetupRouter()
	e.Logger.Fatal(e.Start(":8200"))
}

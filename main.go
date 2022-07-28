package main

import (
	handler "crud/handlers"
)

func main() {

	e := handler.SetupRouter()
	e.Logger.Fatal(e.Start(":1323"))
}

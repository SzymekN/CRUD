package main

import (
	"net/http"

	handler "crud/handlers"

	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/api/v1/users/:id", handler.GetUserById)
	e.GET("/api/v1/users", handler.GetUsers)
	e.POST("/api/v1/users/save", handler.SaveUser)
	e.PUT("/api/v1/users/:id", handler.UpdateUser)
	e.DELETE("/api/v1/users/:id", handler.DeleteUser)

	e.Logger.Fatal(e.Start(":1323"))
}

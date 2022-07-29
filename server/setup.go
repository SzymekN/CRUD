package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func SetupRouter() *echo.Echo {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, `{"message":"Hello World!"}`)
	})
	e.GET("/api/v1/users/:id", GetUserById)
	e.GET("/api/v1/users", GetUsers)
	e.POST("/api/v1/users/save", SaveUser)
	e.PUT("/api/v1/users/:id", UpdateUser)
	e.DELETE("/api/v1/users/:id", DeleteUser)

	return e
}

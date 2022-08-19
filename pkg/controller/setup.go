package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mvrilo/go-redoc"
	echoredoc "github.com/mvrilo/go-redoc/echo"
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

	e.GET("/api/v2/users/:id", GetUserByIdCassandraHandler)
	e.GET("/api/v2/users", GetUsersCassandraHandler)
	e.POST("/api/v2/users/save", SaveUserCassandraHandler)
	e.PUT("/api/v2/users/:id", UpdateUserCassandraHandler)
	e.DELETE("/api/v2/users/:id", DeleteUserCassandraHandler)

	// redoc documentation middleware
	doc := redoc.Redoc{
		Title:       "User API",
		Description: "API for interactions with database",
		SpecFile:    "docs/swagger.json",
		SpecPath:    "docs/swagger.json",
		DocsPath:    "/docs",
	}

	e.Use(echoredoc.New(doc))

	return e
}

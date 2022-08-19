package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mvrilo/go-redoc"
	echoredoc "github.com/mvrilo/go-redoc/echo"
)

func SetupRouter() *echo.Echo {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, `{"message":"Hello World!"}`)
	})

	e.POST("/api/v2/operators/signup", SignUp)
	e.GET("/api/v2/operators/signin", SignIn)

	auth := e.Group("")
	auth.Use(middleware.JWT([]byte(Secretkey)))

	auth.GET("/api/v1/users/:id", GetUserById)
	auth.GET("/api/v1/users", GetUsers)
	auth.POST("/api/v1/users/save", SaveUser)
	auth.PUT("/api/v1/users/:id", UpdateUser, isAdmin)
	auth.DELETE("/api/v1/users/:id", DeleteUser, isAdmin)

	auth.GET("/api/v2/users/:id", GetUserByIdCassandraHandler)
	auth.GET("/api/v2/users", GetUsersCassandraHandler)
	auth.POST("/api/v2/users/save", SaveUserCassandraHandler, isAdmin)
	auth.PUT("/api/v2/users/:id", UpdateUserCassandraHandler, isAdmin)
	auth.DELETE("/api/v2/users/:id", DeleteUserCassandraHandler, isAdmin)

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

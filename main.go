package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type User struct {
	Id         int    `json:"id" query:"id" form:"id" param:"id"`
	First_name string `json:"firstname" query:"firstname" form:"firstname" param:"firstname"`
	Last_name  string `json:"lastname" query:"lastname" form:"lastname" param:"lastname"`
	Age        int    `json:"age" query:"age" form:"age" param:"age"`
}

func updateUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	_, ok := users[id]

	if !ok {
		return c.String(http.StatusNotFound, "User with given id doesn't exist")
	}

	c.Bind(users[id])
	return c.JSON(http.StatusOK, users[id])
}

func saveUser(c echo.Context) error {
	var u User
	if err := c.Bind(u); err != nil {
		return err
	}
	maxID++
	u.Id = maxID
	users[maxID] = &u
	fmt.Println(u)
	return c.JSON(http.StatusOK, users[maxID])
}

// e.GET("/api/v1/users/:id", getUser)
func getUserById(c echo.Context) error {
	// User ID from path `users/:id`
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, users[id])
}

// e.GET("/api/v1/users", getUsers)
func getUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, users)
}

var users = map[int]*User{
	1: {Id: 1, First_name: "Szymon", Last_name: "Nowak", Age: 22},
	2: {Id: 2, First_name: "Jan", Last_name: "Kowalski", Age: 31},
	3: {Id: 3, First_name: "Chuck", Last_name: "Norris", Age: 18},
	4: {Id: 4, First_name: "Andrzej", Last_name: "Duda", Age: 41},
}
var maxID int = 4

func main() {

	fmt.Println(users)

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/api/v1/users/:id", getUserById)
	e.GET("/api/v1/users", getUsers)
	e.POST("/api/v1/users/save", saveUser)
	e.PUT("/api/v1/users/:id", updateUser)

	// e.DELETE("/api/v1/users/:id", deleteUser)

	e.Logger.Fatal(e.Start(":1323"))
}

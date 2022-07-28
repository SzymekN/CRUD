package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

//	e.POST("/api/v1/users/save", saveUser)
func SaveUser(c echo.Context) error {
	var u User
	if err := c.Bind(&u); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	maxID++
	u.Id = maxID
	users[maxID] = &u
	fmt.Println(u)
	return c.JSON(http.StatusOK, users[maxID])
}

//e.PUT("/api/v1/users/:id", updateUser)
func UpdateUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	_, ok := users[id]

	if !ok {
		return c.JSON(http.StatusBadRequest, "User with given id doesn't exist")
	}

	if err := c.Bind(users[id]); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, users[id])
}

//e.DELETE("/api/v1/users/:id", deleteUser)
func DeleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	_, ok := users[id]

	if !ok {
		return c.JSON(http.StatusBadRequest, "User with given id doesn't exist")
	}

	delete(users, id)
	return c.JSON(http.StatusOK, users[id])
}

// e.GET("/api/v1/users/:id", getUser)
func GetUserById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	user, ok := users[id]

	if !ok {
		return c.JSON(http.StatusBadRequest, "User with given id doesn't exist")
	}

	return c.JSON(http.StatusOK, user)
}

// e.GET("/api/v1/users", getUsers)
func GetUsers(c echo.Context) error {
	users_slice := []*User{}

	for _, u := range users {
		users_slice = append(users_slice, u)
	}

	return c.JSON(http.StatusOK, users_slice)
}

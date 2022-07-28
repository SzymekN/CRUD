package handlers

import (
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
	users = append(users, &u)
	return c.JSON(http.StatusOK, users[maxID-1])
}

//e.PUT("/api/v1/users/:id", updateUser)
func UpdateUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	i, ok := findUser(id)

	if !ok {
		return c.JSON(http.StatusBadRequest, `"error":"User with given id doesn't exist"`)
	}

	if err := c.Bind(users[i]); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, users[i])
}

//e.DELETE("/api/v1/users/:id", deleteUser)
func DeleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	i, ok := findUser(id)

	if !ok {
		return c.JSON(http.StatusBadRequest, `"error":"User with given id doesn't exist"`)
	}

	users = removeUser(users, id)
	return c.JSON(http.StatusOK, users[i])
}

// e.GET("/api/v1/users/:id", getUser)
func GetUserById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	i, ok := findUser(id)

	if !ok {
		return c.JSON(http.StatusBadRequest, `"error":"User with given id doesn't exist"`)
	}

	return c.JSON(http.StatusOK, users[i])
}

// e.GET("/api/v1/users", getUsers)
func GetUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, users)
}

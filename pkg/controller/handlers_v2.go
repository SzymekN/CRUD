package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"crud/pkg/model"

	"github.com/labstack/echo/v4"
)

//	e.POST("/api/v2/users/save", saveUserCassandra)
func SaveUserCassandraHandler(c echo.Context) error {
	var u model.User
	if err := c.Bind(&u); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	SaveUserCassandra(u)
	return c.JSON(http.StatusOK, u)
}

//	e.PUT("/api/v2/users/:id", updateUser)
func UpdateUserCassandraHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.String(http.StatusBadRequest, `"error":"id not valid"`)
	}

	u, err := GetUserByIdCassandra(id)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, `"error":"user doesn't exits"`)
	}

	if err := c.Bind(&u); err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, `"error":"couldn't update user"`)
	}

	err = UpdateUserCassandra(id, u)

	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, `"error":"couldn't update user"`)
	}

	return c.JSON(http.StatusOK, u)
}

//	e.DELETE("/api/v2/users/:id", deleteUser)
func DeleteUserCassandraHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.String(http.StatusBadRequest, `"error":"user with given id doesn't exist"`)
	}

	_, err = DeleteUserCassandra(id)

	if err != nil {
		return c.String(http.StatusOK, `"message":"error while deleting user"`)
	}

	return c.String(http.StatusOK, `"message":"user deleted"`)
}

// e.GET("/api/v2/users/:id", getUser)
func GetUserByIdCassandraHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.String(http.StatusBadRequest, `"error":"user with given id doesn't exist"`)
	}

	u, err := GetUserByIdCassandra(id)
	if err != nil {
		return c.String(http.StatusBadRequest, `"error":"couldn't get user"`)
	}

	return c.JSON(http.StatusOK, u)
}

// e.GET("/api/v2/users", getUsers)
func GetUsersCassandraHandler(c echo.Context) error {

	users, err := GetUsersCassandra()
	if err != nil {
		return c.String(http.StatusBadRequest, `"error":"couldn't get users"`)
	}
	return c.JSON(http.StatusOK, users)
}

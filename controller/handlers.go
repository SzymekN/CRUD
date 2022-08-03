package controller

import (
	"net/http"
	"strconv"

	"crud/model"
	storage "crud/storage"

	"github.com/labstack/echo/v4"
)

//	e.POST("/api/v1/users/save", saveUser)
func SaveUser(c echo.Context) error {
	var u model.User
	if err := c.Bind(&u); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	db := storage.GetDBInstance()
	db.Create(&u)
	return c.JSON(http.StatusOK, u)
}

//	e.PUT("/api/v1/users/:id", updateUser)
func UpdateUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.String(http.StatusBadRequest, `"error":"id not valid"`)
	}

	db := storage.GetDBInstance()
	user := model.User{}
	result := db.Find(&user, id)

	if result.RowsAffected < 1 {
		return c.String(http.StatusBadRequest, `"error":"user not found"`)
	}

	if err := c.Bind(&user); err != nil {
		return c.String(http.StatusBadRequest, `"error":"couldn't update user"`)
	}

	user.Id = id
	result = db.Save(&user)
	if result.RowsAffected < 1 {
		return c.String(http.StatusBadRequest, `"error":"couldn't update user"`)
	}

	return c.JSON(http.StatusOK, user)
}

//	e.DELETE("/api/v1/users/:id", deleteUser)
func DeleteUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.String(http.StatusBadRequest, `"error":"user with given id doesn't exist"`)
	}

	db := storage.GetDBInstance()
	result := db.Delete(&model.User{}, id)

	if result.RowsAffected < 1 {
		return c.String(http.StatusBadRequest, `"error":"user with given id doesn't exist"`)
	}

	return c.String(http.StatusOK, `"message":"user deleted"`)
}

// e.GET("/api/v1/users/:id", getUser)
func GetUserById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.String(http.StatusBadRequest, `"error":"user with given id doesn't exist"`)
	}

	db := storage.GetDBInstance()
	user := model.User{}
	result := db.Find(&user, id)

	if result.RowsAffected < 1 {
		return c.String(http.StatusBadRequest, `"error":"user with given id doesn't exist"`)
	}

	return c.JSON(http.StatusOK, user)
}

// e.GET("/api/v1/users", getUsers)
func GetUsers(c echo.Context) error {
	db := storage.GetDBInstance()
	users := []model.User{}

	if err := db.Find(&users).Error; err != nil {
		return nil
	}

	return c.JSON(http.StatusOK, users)
}

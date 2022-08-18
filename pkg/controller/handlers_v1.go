package controller

import (
	"net/http"
	"strconv"

	"github.com/SzymekN/CRUD/pkg/model"
	"github.com/SzymekN/CRUD/pkg/storage"

	"github.com/labstack/echo/v4"
)

// swagger:route POST /api/v1/users/save users_v1 postUserV1
// Save user to postgres database.
//	Consumes:
//    - application/json
//  Produces:
//    - application/json
//
// responses:
// 		200: userResponse
//		500: errorResponse
func SaveUser(c echo.Context) error {
	var u model.User
	if err := c.Bind(&u); err != nil {
		return c.JSON(http.StatusInternalServerError, &model.GenericError{Message: "couldn't update users"})
	}
	db := storage.GetDBInstance()
	db.Create(&u)
	return c.JSON(http.StatusOK, u)
}

// swagger:route PUT /api/v1/user/{id} users_v1 putUserV1
// Updates user in postgres database.
//	Consumes:
//    - application/json
//  Produces:
//    - application/json
//
// responses:
// 		200: userResponse
//		400: errorResponse
//		404: errorResponse
//		500: errorResponse
func UpdateUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, &model.GenericError{Message: `"error":"incorrect id"`})
	}

	db := storage.GetDBInstance()
	user := model.User{}
	result := db.Find(&user, id)

	if result.RowsAffected < 1 {
		return c.JSON(http.StatusNotFound, &model.GenericError{Message: "user doesn't exist"})
	}

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusInternalServerError, &model.GenericError{Message: "couldn't update users"})
	}

	user.Id = id
	result = db.Save(&user)
	if result.RowsAffected < 1 {
		return c.JSON(http.StatusInternalServerError, &model.GenericError{Message: "couldn't update users"})
	}

	return c.JSON(http.StatusOK, user)
}

// swagger:route DELETE /api/v1/user/{id} users_v1 deleteUserV1
// deletes user from postgres database.
//  Produces:
//    - application/json
//
// responses:
// 		200: messageResponse
//		400: errorResponse
//		404: errorResponse
func DeleteUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, &model.GenericError{Message: `"error":"incorrect id"`})
	}

	db := storage.GetDBInstance()
	result := db.Delete(&model.User{}, id)

	if result.RowsAffected < 1 {
		return c.JSON(http.StatusNotFound, &model.GenericError{Message: "couldn't get user"})
	}

	return c.JSON(http.StatusOK, &model.GenericMessage{Message: "user deleted"})
}

// swagger:route GET /api/v1/user/{id} users_v1 getUserV1
// Gets user from postgres database.
//  Produces:
//    - application/json
//
// responses:
// 		200: userResponse
//		400: errorResponse
//		404: errorResponse
func GetUserById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, &model.GenericError{Message: `"error":"incorrect id"`})
	}

	db := storage.GetDBInstance()
	user := model.User{}
	result := db.Find(&user, id)

	if result.RowsAffected < 1 {
		return c.JSON(http.StatusNotFound, &model.GenericError{Message: "couldn't get user"})
	}

	return c.JSON(http.StatusOK, user)
}

// swagger:route GET /api/v1/users users_v1 listUsersV1
// Gets user from postgres database.
//
//  Produces:
//    - application/json
//
// responses:
// 		200: usersResponse
//		500: errorResponse
func GetUsers(c echo.Context) error {
	db := storage.GetDBInstance()
	users := []model.User{}

	if err := db.Find(&users).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, &model.GenericError{Message: "couldn't get users"})
	}

	return c.JSON(http.StatusOK, users)
}

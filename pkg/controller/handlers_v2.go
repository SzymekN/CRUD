package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/SzymekN/CRUD/pkg/model"

	"github.com/labstack/echo/v4"
)

// swagger:route POST /api/v2/users/save users_v2 postUserV2
// Save user to cassandra database.
//	Consumes:
//    - application/json
//  Produces:
//    - application/json
//
// responses:
// 		200: userResponse
//		500: errorResponse
func SaveUserCassandraHandler(c echo.Context) error {
	var u model.User
	if err := c.Bind(&u); err != nil {
		return c.JSON(http.StatusInternalServerError, &model.GenericError{Message: "couldn't update users"})
	}
	SaveUserCassandra(u)
	return c.JSON(http.StatusOK, u)
}

// swagger:route PUT /api/v2/user/{id} users_v2 putUserV2
// Updates user in cassandra database.
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
func UpdateUserCassandraHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, &model.GenericError{Message: `"error":"incorrect id"`})
	}

	u, err := GetUserByIdCassandra(id)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusNotFound, &model.GenericError{Message: "user doesn't exist"})
	}

	if err := c.Bind(&u); err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, &model.GenericError{Message: "couldn't update users"})
	}

	err = UpdateUserCassandra(id, u)

	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, &model.GenericError{Message: "couldn't update users"})
	}

	return c.JSON(http.StatusOK, u)
}

// swagger:route DELETE /api/v2/user/{id} users_v2 deleteUserV2
// deletes user from cassandra database.
//  Produces:
//    - application/json
//
// responses:
// 		200: messageResponse
//		400: errorResponse
//		404: errorResponse
func DeleteUserCassandraHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, &model.GenericError{Message: `"error":"incorrect id"`})
	}

	_, err = DeleteUserCassandra(id)

	if err != nil {
		return c.JSON(http.StatusNotFound, &model.GenericError{Message: "couldn't get user"})
	}

	return c.JSON(http.StatusOK, &model.GenericMessage{Message: "user deleted"})
}

// swagger:route GET /api/v2/user/{id} users_v2 getUserV2
// Gets user from cassandra database.
//  Produces:
//    - application/json
//
// responses:
// 		200: userResponse
//		400: errorResponse
//		404: errorResponse
func GetUserByIdCassandraHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, &model.GenericError{Message: `"error":"incorrect id"`})
	}

	u, err := GetUserByIdCassandra(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, &model.GenericError{Message: "couldn't get user"})
	}

	return c.JSON(http.StatusOK, u)
}

// swagger:route GET /api/v2/users users_v2 getUsersV2
// Gets user from cassandra database.
//
//  Produces:
//    - application/json
//
// responses:
// 		200: usersResponse
//		500: errorResponse
func GetUsersCassandraHandler(c echo.Context) error {

	users, err := GetUsersCassandra()
	if err != nil {
		return c.String(http.StatusBadRequest, `"error":"couldn't get users"`)
	}
	return c.JSON(http.StatusOK, users)
}

package controller

import (
	"net/http"
	"strconv"

	"github.com/SzymekN/CRUD/pkg/model"
	"github.com/SzymekN/CRUD/pkg/producer"

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
	var err error
	var status int
	k, msg := "", "userapi.users"

	defer func() {
		producer.ProduceMessage(k, msg)
		if err != nil {
			c.JSON(status, &model.GenericError{Message: msg})
		}
	}()

	if err := c.Bind(&u); err != nil {
		status = http.StatusBadRequest
		msg += "[" + k + "] POST error: incorrect parameters, HTTP: " + strconv.Itoa(status)
		return err
	}

	k = strconv.Itoa(u.Id)
	err = SaveUserCassandra(u)
	if err != nil {
		status = http.StatusInternalServerError
		msg += "[" + k + "] POST error: post query error, HTTP: " + strconv.Itoa(status)
		return err
	}

	status = http.StatusOK
	msg += "[" + k + "] POST completed: user added, HTTP: " + strconv.Itoa(status)
	c.JSON(http.StatusOK, u)
	return nil
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
	k, msg := "", "userapi.users"

	var status int

	defer func() {
		producer.ProduceMessage(k, msg)
		if err != nil {
			c.JSON(status, &model.GenericError{Message: msg})
		}
	}()

	if err != nil {
		k = "unknown"
		status = http.StatusBadRequest
		msg += "[" + k + "] PUT error: incorrect id, HTTP: " + strconv.Itoa(status)
		return err
	}

	k = strconv.Itoa(id)
	u, err := GetUserByIdCassandra(id)
	if err != nil {
		status = http.StatusNotFound
		msg += "[" + k + "] PUT error: user doesn't exist, HTTP: " + strconv.Itoa(status)
		return err
	}

	if err := c.Bind(&u); err != nil {
		status = http.StatusBadRequest
		msg += "[" + k + "] PUT error: incorrect parameters, HTTP: " + strconv.Itoa(status)
		return err
	}

	err = UpdateUserCassandra(id, u)

	if err != nil {
		status = http.StatusInternalServerError
		msg += "[" + k + "] PUT error: update query error, HTTP: " + strconv.Itoa(status)
		return err
	}

	status = http.StatusOK
	msg += "[" + k + "] PUT completed: user updated, HTTP: " + strconv.Itoa(status)
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
	k, msg := "", "userapi.users"

	var status int

	defer func() {
		producer.ProduceMessage(k, msg)
		if err != nil {
			c.JSON(status, &model.GenericError{Message: msg})
		}
	}()

	if err != nil {
		k = "unknown"
		status = http.StatusBadRequest
		msg += "[" + k + "] DEL error: incorrect id, HTTP: " + strconv.Itoa(status)
		return err
	}

	_, err = DeleteUserCassandra(id)
	k = strconv.Itoa(id)

	if err != nil {
		status = http.StatusNotFound
		msg += "[" + k + "] DEL error: user doesn't exist, HTTP: " + strconv.Itoa(status)
		return err
	}

	status = http.StatusOK
	msg += "[" + k + "] DEL completed: user deleted, HTTP: " + strconv.Itoa(status)
	return c.JSON(status, &model.GenericMessage{Message: msg})
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

	k, msg := "", "userapi.users"

	var status int

	defer func() {
		producer.ProduceMessage(k, msg)
		if err != nil {
			c.JSON(status, &model.GenericError{Message: msg})
		}
	}()

	if err != nil {
		k = "unknown"
		status = http.StatusBadRequest
		msg += "[" + k + "] GET error: incorrect id, HTTP: " + strconv.Itoa(status)
		return err
	}

	k = strconv.Itoa(id)
	u, err := GetUserByIdCassandra(id)
	if err != nil {
		status = http.StatusNotFound
		msg += "[" + k + "] GET error: couldn't get user, HTTP: " + strconv.Itoa(status)
		return err
	}

	status = http.StatusOK
	msg += "[" + k + "] GET completed: user read, HTTP: " + strconv.Itoa(status)
	return c.JSON(status, u)
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

	k, msg := "", "userapi.users"

	var status int

	defer func() {
		producer.ProduceMessage(k, msg)
	}()

	k = "all"
	if err != nil {
		status = http.StatusNotFound
		msg += "[" + k + "] GET error: couldn't get users, HTTP: " + strconv.Itoa(status)
		return err
	}

	status = http.StatusOK
	msg += "[" + k + "] GET completed: users read, HTTP: " + strconv.Itoa(status)
	return c.JSON(http.StatusOK, users)
}

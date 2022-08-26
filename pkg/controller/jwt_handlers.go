package controller

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/SzymekN/CRUD/pkg/auth"
	"github.com/SzymekN/CRUD/pkg/model"
	"github.com/SzymekN/CRUD/pkg/producer"
	"github.com/SzymekN/CRUD/pkg/storage"
	"github.com/gocql/gocql"
	"github.com/labstack/echo/v4"
)

func GetOperator(username string) (auth.Operator, error) {
	cas := storage.GetCassandraInstance()
	o := auth.Operator{}

	if err := cas.Query(`Select username, email, password, role from userapi.operators where username=?`, username).Consistency(gocql.One).Scan(&o.Username, &o.Email, &o.Password, &o.Role); err != nil {
		fmt.Println(err)
		return o, err
	}

	return o, nil
}

func SaveOperator(op auth.Operator) error {
	cas := storage.GetCassandraInstance()
	if err := cas.Query(`Insert into userapi.operators(username, email,password, role) values (?,?,?,?)`, op.Username, op.Email, op.Password, op.Role).Exec(); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func SignUp(c echo.Context) error {

	var op auth.Operator
	var err error
	var status int
	k, msg := "", "userapi.operators"

	defer func() {
		producer.ProduceMessage(k, msg)
		if err != nil {
			c.JSON(status, &model.GenericError{Message: msg})
		}
	}()

	if err = c.Bind(&op); err != nil {
		status = http.StatusBadRequest
		k = op.Username
		msg += "[" + k + "] SignUp error: incorrect credentials, HTTP: " + strconv.Itoa(status)
		return err
	}
	_, err = GetOperator(op.Username)

	k = op.Username
	if err == nil {
		status = http.StatusInternalServerError
		msg += "[" + k + "] SignUp error: username in use, HTTP: " + strconv.Itoa(status)
		err = errors.New("user exists")
		return err
	}

	op.Password, err = auth.GeneratehashPassword(op.Password)
	if err != nil {
		status = http.StatusInternalServerError
		msg += "[" + k + "] SignUp error: couldn't generate hash, HTTP: " + strconv.Itoa(status)
		return err
	}

	//insert user details in database
	err = SaveOperator(op)
	if err != nil {
		status = http.StatusInternalServerError
		msg += "[" + k + "] SignUp error: insert query error, HTTP: " + strconv.Itoa(status)
		return err
	}

	status = http.StatusOK
	msg += "[" + k + "] SignUp completed: user signed up, HTTP: " + strconv.Itoa(status)
	return c.JSON(http.StatusOK, op)

}

func SignIn(c echo.Context) error {

	var authDetails auth.Authentication
	var err error
	var status int
	k, msg := "", "userapi.operators"

	defer func() {
		producer.ProduceMessage(k, msg)
		if err != nil {
			c.JSON(status, &model.GenericError{Message: msg})
		}
	}()

	if err = c.Bind(&authDetails); err != nil {
		status = http.StatusBadRequest
		k = authDetails.Username
		msg += "[" + k + "] SignIn error: incorrect credentials, HTTP: " + strconv.Itoa(status)
		return err
	}

	var authUser auth.Operator
	authUser, err = GetOperator(authDetails.Username)

	k = authDetails.Username
	if err != nil {
		status = http.StatusInternalServerError
		msg += "[" + k + "] SignIn error: user doesn't exist, HTTP: " + strconv.Itoa(status)
		return err
	}

	check := auth.CheckPasswordHash(authDetails.Password, authUser.Password)

	if !check {
		status = http.StatusBadRequest
		msg += "[" + k + "] SignIn error: incorrect password, HTTP: " + strconv.Itoa(status)
		err = errors.New("Incorrect password")
		return err
	}

	var validToken string
	validToken, err = auth.GenerateJWT(authDetails.Username, authUser.Role)
	if err != nil {
		status = http.StatusInternalServerError
		msg += "[" + k + "] SignIn error: couldn't generate token, HTTP: " + strconv.Itoa(status)
		return err
	}

	var token auth.Token
	token.Username = authUser.Username
	token.Role = authUser.Role
	token.TokenString = validToken
	status = http.StatusOK
	msg += "[" + k + "] SignIn completed: user signed in, HTTP: " + strconv.Itoa(status)
	return c.JSON(status, token)
}

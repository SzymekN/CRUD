package controller

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"crud/pkg/model"
	"crud/pkg/storage"

	"github.com/gocql/gocql"
	"github.com/labstack/echo/v4"
)

//	e.POST("/api/v2/users/save", saveUserCassandra)
func SaveUserCassandra(c echo.Context) error {
	var u model.User
	if err := c.Bind(&u); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	cas := storage.GetCassandraInstance()
	if err := cas.Query(`Insert into userapi.users(id, firstname,lastname, age) values (?,?,?,?)`, u.Id, u.Firstname, u.Lastname, u.Age).Exec(); err != nil {
		fmt.Println(err)
	}
	return c.JSON(http.StatusOK, u)
}

//	e.PUT("/api/v2/users/:id", updateUser)
func UpdateUserCassandra(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.String(http.StatusBadRequest, `"error":"id not valid"`)
	}

	cas := storage.GetCassandraInstance()
	u := model.User{}
	if err := cas.Query(`Select id, firstname, lastname, age from userapi.users where id=?`, id).Consistency(gocql.One).Scan(&u.Id, &u.Firstname, &u.Lastname, &u.Age); err != nil {
		fmt.Println(err)
	}

	if err := c.Bind(&u); err != nil {
		return c.String(http.StatusBadRequest, `"error":"couldn't update user"`)
	}

	if err := cas.Query(`Update userapi.users set firstname=?, lastname=?, age=? where id=?`, u.Firstname, u.Lastname, u.Age, id).Exec(); err != nil {
		fmt.Println(err)
	}
	return c.JSON(http.StatusOK, u)
}

//	e.DELETE("/api/v2/users/:id", deleteUser)
func DeleteUserCassandra(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.String(http.StatusBadRequest, `"error":"user with given id doesn't exist"`)
	}

	cas := storage.GetCassandraInstance()
	if err := cas.Query(`Delete from userapi.users where id=?`, id).Exec(); err != nil {
		fmt.Println(err)
	}

	return c.String(http.StatusOK, `"message":"user deleted"`)
}

// e.GET("/api/v2/users/:id", getUser)
func GetUserByIdCassandra(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.String(http.StatusBadRequest, `"error":"user with given id doesn't exist"`)
	}

	cas := storage.GetCassandraInstance()
	u := model.User{}

	if err := cas.Query(`Select id, firstname, lastname, age from userapi.users where id=?`, id).Consistency(gocql.One).Scan(&u.Id, &u.Firstname, &u.Lastname, &u.Age); err != nil {
		fmt.Println(err)
	}

	return c.JSON(http.StatusOK, u)
}

// e.GET("/api/v2/users", getUsers)
func GetUsersCassandra(c echo.Context) error {
	users := []model.User{}
	u := model.User{}

	cas := storage.GetCassandraInstance()

	iter := cas.Query(`Select id, firstname, lastname, age from userapi.users`).Iter()
	for iter.Scan(&u.Id, &u.Firstname, &u.Lastname, &u.Age) {
		users = append(users, u)
		u = model.User{}
	}
	if err := iter.Close(); err != nil {
		log.Fatal(err)
	}

	return c.JSON(http.StatusOK, users)
}

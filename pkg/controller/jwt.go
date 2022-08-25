package controller

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/SzymekN/CRUD/pkg/model"
	"github.com/SzymekN/CRUD/pkg/storage"
	"github.com/gocql/gocql"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

var Secretkey string = "tokluczjestjakis"

type Operator struct {
	Username string `json:"username" form:"username"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Role     string `json:"role" form:"role"`
}

type Authentication struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

type Token struct {
	Role        string `json:"role" form:"role"`
	Username    string `json:"username" form:"username"`
	TokenString string `json:"token" form:"token"`
}

func GetOperator(username string) (Operator, error) {
	cas := storage.GetCassandraInstance()
	o := Operator{}

	if err := cas.Query(`Select username, email, password, role from userapi.operators where username=?`, username).Consistency(gocql.One).Scan(&o.Username, &o.Email, &o.Password, &o.Role); err != nil {
		fmt.Println(err)
		return o, err
	}

	return o, nil
}

func SaveOperator(op Operator) error {
	cas := storage.GetCassandraInstance()
	if err := cas.Query(`Insert into userapi.operators(username, email,password, role) values (?,?,?,?)`, op.Username, op.Email, op.Password, op.Role).Exec(); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func SignUp(c echo.Context) error {

	var op Operator

	if err := c.Bind(&op); err != nil {
		return c.JSON(http.StatusInternalServerError, &model.GenericError{Message: "incorrect credantials"})
	}
	_, err := GetOperator(op.Username)

	if err == nil {
		return c.JSON(http.StatusInternalServerError, &model.GenericError{Message: "username in use"})
	}

	op.Password, err = GeneratehashPassword(op.Password)
	if err != nil {
		log.Fatalln("error in password hash")
	}

	//insert user details in database
	SaveOperator(op)
	return c.JSON(http.StatusOK, op)

}

func SignIn(c echo.Context) error {

	var authDetails Authentication

	if err := c.Bind(&authDetails); err != nil {
		return c.JSON(http.StatusInternalServerError, &model.GenericError{Message: "incorrect credantials"})
	}

	authUser, err := GetOperator(authDetails.Username)

	if err != nil {
		return c.JSON(http.StatusNotFound, &model.GenericError{Message: "user doesn't exist"})
	}

	check := CheckPasswordHash(authDetails.Password, authUser.Password)

	if !check {
		return c.JSON(http.StatusBadRequest, &model.GenericError{Message: "incorrect password"})

	}

	validToken, err := GenerateJWT(authDetails.Username, authUser.Role)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &model.GenericError{Message: "couldn't generate token"})
	}

	var token Token
	token.Username = authUser.Username
	token.Role = authUser.Role
	token.TokenString = validToken
	return c.JSON(http.StatusOK, token)
}

func GeneratehashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateJWT(username, role string) (string, error) {
	var mySigningKey = []byte(Secretkey)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["username"] = username
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Errorf("Something Went Wrong: %s", err.Error())
		return "", err
	}
	return tokenString, nil
}

func isAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		role := claims["role"]
		if role != "admin" {
			return echo.ErrUnauthorized
		}
		return next(c)
	}
}

// func IsAuthorized(handler http.HandlerFunc) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {

// 		if r.Header["Token"] == nil {
// 			var err Error
// 			err = SetError(err, "No Token Found")
// 			json.NewEncoder(w).Encode(err)
// 			return
// 		}

// 		var mySigningKey = []byte(Secretkey)

// 		token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
// 			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 				return nil, fmt.Errorf("There was an error in parsing")
// 			}
// 			return mySigningKey, nil
// 		})

// 		if err != nil {
// 			var err Error
// 			err = SetError(err, "Your Token has been expired")
// 			json.NewEncoder(w).Encode(err)
// 			return
// 		}

// 		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
// 			if claims["role"] == "admin" {

// 				r.Header.Set("Role", "admin")
// 				handler.ServeHTTP(w, r)
// 				return

// 			} else if claims["role"] == "user" {

// 				r.Header.Set("Role", "user")
// 				handler.ServeHTTP(w, r)
// 				return
// 			}
// 		}
// 		var reserr Error
// 		reserr = SetError(reserr, "Not Authorized")
// 		json.NewEncoder(w).Encode(err)
// 	}
// }

package auth

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/SzymekN/CRUD/pkg/producer"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

var Secretkey string = os.Getenv("JWT_KEY")

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

func GeneratehashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func Validate(auth string, c echo.Context) (interface{}, error) {

	localKeyFunc := func(t *jwt.Token) (interface{}, error) {
		if t.Method.Alg() != "HS256" {
			return nil, fmt.Errorf("unexpected jwt signing method=%v", t.Header["alg"])
		}
		return []byte(Secretkey), nil
	}

	remoteKeyFunc := func(t *jwt.Token) (interface{}, error) {
		if t.Method.Alg() != "HS256" {
			return nil, fmt.Errorf("unexpected jwt signing method=%v", t.Header["alg"])
		}
		key, _ := getSigningKey(t.Raw)
		return []byte(key), nil
	}

	// claims are of type `jwt.MapClaims` when token is created with `jwt.Parse`
	token, err := jwt.Parse(auth, localKeyFunc)

	// try parsing with remote key
	if err != nil {
		token, err = jwt.Parse(auth, remoteKeyFunc)
	}

	// zwrócony token i nil == poprawny token
	if err != nil {
		producer.ProduceMessage("JWT validation", "JWT validation failed: "+err.Error())
		return nil, err
	}
	if !token.Valid {
		producer.ProduceMessage("JWT validation", "JWT validation failed: "+err.Error())
		return nil, errors.New("invalid token")
	}

	producer.ProduceMessage("JWT validation", "JWT validation succesfull")
	return token, nil
}

func GenerateJWT(username, role string) (string, error) {
	var mySigningKey = []byte(Secretkey)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	expireTime := time.Minute * 2

	claims["authorized"] = true
	claims["username"] = username
	claims["role"] = role
	claims["exp"] = time.Now().Add(expireTime).Unix()

	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		fmt.Errorf("Something Went Wrong: %s", err.Error())
		return "", err
	}
	// SetToken(tokenString, Secretkey, expireTime)
	SetToken(tokenString, Secretkey, expireTime)
	return tokenString, nil
}

func IsAdmin(next echo.HandlerFunc) echo.HandlerFunc {
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

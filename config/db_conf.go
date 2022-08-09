package config

import (
	"fmt"
	"os"
)

var (
	DBUser     = "userapi"
	DBPassword = "userapi"
	DBName     = "userapi"
	DBHost     = "192.168.33.50"
	DBPort     = "5432"
	DBType     = "postgres"
)

func GetDBType() string {
	return DBType
}

func readEnv() {
	if os.Getenv("DB_USER") != "" {
		DBUser = os.Getenv("DB_USER")
	}

	if os.Getenv("DB_PASSWORD") != "" {
		DBPassword = os.Getenv("DB_PASSWORD")
	}

	if os.Getenv("DB_NAME") != "" {
		DBName = os.Getenv("DB_NAME")
	}

	// if os.Getenv("DB_HOST") != "" {
	// 	DBHost = os.Getenv("DB_HOST")
	// }

	if os.Getenv("DB_PORT") != "" {
		DBPort = os.Getenv("DB_PORT")
	}
}

func GetPostgresConnectionString() string {

	readEnv()
	dataBase := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		DBHost,
		DBPort,
		DBUser,
		DBName,
		DBPassword)
	return dataBase
}

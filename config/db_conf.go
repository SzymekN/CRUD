package config

import "fmt"

const (
	DBUser     = "userapi"
	DBPassword = "userapi"
	DBName     = "userapi"
	DBHost     = "192.168.33.30"
	DBPort     = "5432"
	DBType     = "postgres"
)

func GetDBType() string {
	return DBType
}

func GetPostgresConnectionString() string {
	dataBase := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		DBHost,
		DBPort,
		DBUser,
		DBName,
		DBPassword)
	return dataBase
}

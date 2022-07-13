package helper

import (
	"simple-go-app/src/infrastructure/config"
	"simple-go-app/src/infrastructure/db"
	"strconv"
)

func CreateMySQLConfig(config config.Configuration) *db.MySQLConfig {
	connMaxLifetime, _ := strconv.Atoi(config.Get("MYSQL_CON_MAX_LIFE_TIME"))
	maxOpenConns, _ := strconv.Atoi(config.Get("MYSQL_MAX_OPEN_CON"))
	maxIdleConns, _ := strconv.Atoi(config.Get("MYSQL_MAX_IDLE_CON"))

	return &db.MySQLConfig{
		Host:            config.Get("MYSQL_HOST"),
		Port:            config.Get("MYSQL_PORT"),
		User:            config.Get("MYSQL_USERNAME"),
		Password:        config.Get("MYSQL_PASSWORD"),
		Database:        config.Get("MYSQL_DATABASE"),
		ConnMaxLifetime: connMaxLifetime,
		MaxOpenConns:    maxOpenConns,
		MaxIdleConns:    maxIdleConns,
	}
}

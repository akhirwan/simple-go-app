package db

import (
	"fmt"
	"log"
	"simple-procurement/src/infrastructure/config"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type MySQLConfig struct {
	Host            string
	Port            string
	User            string
	Password        string
	Database        string
	ConnMaxLifetime int
	MaxOpenConns    int
	MaxIdleConns    int
}

func NewMySQLDBConnection(c *MySQLConfig, config config.Configuration) *sqlx.DB {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?timeout=3s&charset=utf8mb4&parseTime=true&loc=Local",
		c.User,
		c.Password,
		c.Host,
		c.Port,
		c.Database)

	db, err := sqlx.Open("mysql", dataSourceName)
	if err != nil {
		log.Printf("Connect Database failed: %s", err.Error())
		return nil
	}

	if err := db.Ping(); err != nil {
		db.Close()
		log.Fatalf("Database is unreachable. %s", err.Error())
		return nil
	}

	log.Printf("Successfully connected to the db.\n")

	db.SetConnMaxLifetime(time.Minute * time.Duration(c.ConnMaxLifetime))
	db.SetMaxOpenConns(c.MaxOpenConns)
	db.SetMaxIdleConns(c.MaxIdleConns)

	log.Println("MySQL Database connected")

	return db
}

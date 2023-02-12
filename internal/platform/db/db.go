package db

import (
	"fmt"
	"os"
	"time"

	// postgres db driver
	_ "github.com/lib/pq"

	"github.com/jmoiron/sqlx"
)

const DriverName = "postgres"

func Init() (*sqlx.DB, error) {
	dataSource := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_USERNAME"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DBNAME"))

	client, err := sqlx.Open(DriverName, dataSource)
	if err != nil {
		return nil, err
	}

	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	// verifies connection is db is working
	if err := client.Ping(); err != nil {
		return nil, err
	}

	return client, nil
}

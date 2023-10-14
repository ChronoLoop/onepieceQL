package db

import (
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/pressly/goose/v3"
	"log"
	"os"
)

var Connection *sqlx.DB

func Initialize() {
	host, port, username, password, name :=
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		host,
		username,
		password,
		name,
		port,
	)

	db, err := sqlx.Connect("pgx", dsn)
	if err != nil {
		log.Fatal(err.Error())
	}

	migrationDir := "./internal/server/db/migrations"

	if err = goose.Up(db.DB, migrationDir); err != nil {
		log.Fatal(err.Error())
	}

	Connection = db
}

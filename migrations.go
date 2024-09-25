package main

import (
	"database/sql"
	"embed"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

//go:embed db/migrations/*.sql
var embedMigrations embed.FS

func main() {
	var db *sql.DB
	host := "localhost"
	port := 5432
	user := "admin"
	password := "admin"
	dbName := "clochness"

	var connectString = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)

	db, err := sql.Open("postgres", connectString)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}

	if err := goose.Up(db, "db/migrations"); err != nil {
		panic(err)
	}
}

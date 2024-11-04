package db

import (
	"context"
	"database/sql"
	_ "embed"
	_ "github.com/mattn/go-sqlite3"
)

//go:embed schema.sql
var ddl string

func Provide() *Queries {
	db, err := sql.Open("sqlite3", "file:htmgo-user-example.db?cache=shared&_fk=1")

	if err != nil {
		panic(err)
	}

	if _, err := db.ExecContext(context.Background(), ddl); err != nil {
		panic(err)
	}

	return New(db)
}

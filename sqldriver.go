package main

import (
	"database/sql"
	"fmt"
	"os"
)

func openSql() *sql.DB {
	db, err := sql.Open("sqlite", "database.db")

	if err != nil {
		fmt.Printf("encountered error for connection: %s", err.Error())
		os.Exit(68)
	}

	return db
}

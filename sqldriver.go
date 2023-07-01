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

func executeQuery(db *sql.DB, query string, args ...any) sql.Result {
	var res sql.Result
	var err error
	if args != nil {
		res, err = db.Exec(query, args...)
	} else {
		res, err = db.Exec(query)
	}

	if err != nil {
		fmt.Printf("Encountered error while executing the query.\nQuery: %s\nError:%s\n", query, err.Error())
		os.Exit(69)
	}

	return res
}

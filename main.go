package main

import (
	"database/sql"
	"fmt"
	_ "modernc.org/sqlite"
	"os"
)

const create string = `
  CREATE TABLE IF NOT EXISTS overtime (
    id INTEGER NOT NULL PRIMARY KEY,
    name TEXT NOT NULL,
    time_in TEXT NOT NULL,
    time_out TEXT NOT NULL,
    hours_ot INTEGER NOT NULL,
    reason TEXT NOT NULL
  );
`

func main() {
	args := os.Args[1:]
	ot, err := govertime(args)
	if err != nil {
		fmt.Printf("encountered error: %s", err.Error())
		os.Exit(69)
	}

	db, err := sql.Open("sqlite", "database.db")
	if err != nil {
		fmt.Printf("encountered error for connection: %s", err.Error())
		os.Exit(68)
	}

	if _, err := db.Exec(create); err != nil {
		fmt.Printf("encountered error: %s", err.Error())
		os.Exit(70)
	}

	for _, v := range ot {
		res, err := db.Exec("INSERT INTO overtime VALUES(NULL, ?, ?, ?, ?, ?)", v.Name, v.TimeIn, v.TimeOut, v.HoursOT, v.Reason)
		rows, err := res.RowsAffected()
		if err != nil {
			fmt.Printf("encountered error while inserting: %s", err.Error())
			os.Exit(71)
		}

		fmt.Printf("Inserted %d row! %v\n", rows, v)
	}

	if err := db.Close(); err != nil {
		fmt.Printf("encountered error while closing db: %s", err.Error())
		os.Exit(71)
	}
}

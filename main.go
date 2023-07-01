package main

import (
	"fmt"
	"os"

	_ "modernc.org/sqlite"
)

func main() {
	shouldExit := handleFlags()
	if shouldExit {
		return
	}

	args := os.Args[1:]
	ot, err := govertime(args)
	if err != nil {
		fmt.Printf("encountered error: %s", err.Error())
		os.Exit(69)
	}

	db := openSql()
	defer db.Close()

	executeQuery(db, CreateTableQuery)

	for _, v := range ot {
		res := executeQuery(db, "INSERT INTO overtime VALUES(NULL, ?, ?, ?, ?, ?)", v.Name, v.TimeIn, v.TimeOut, v.HoursOT, v.Reason)
		rows, err := res.RowsAffected()
		if err != nil {
			fmt.Printf("Failed to insert record. Error %s", err.Error())
		}

		fmt.Printf("Inserted %d row! %v\n", rows, v)
	}
}

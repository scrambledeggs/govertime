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

	if _, err := db.Exec(CreateTableQuery); err != nil {
		fmt.Printf("Encountered error while executing the query.\nQuery: %s\nError:%s\n", CreateTableQuery, err.Error())
		os.Exit(SQLQueryError)
	}

	for _, v := range ot {
		res, err := db.Exec(InsertOvertimeQuery, v.Name, v.TimeIn, v.TimeOut, v.HoursOT, v.Reason)
		if err != nil {
			fmt.Printf("Encountered error while executing the query.\nQuery: %s\nError:%s\n", CreateTableQuery, err.Error())
			os.Exit(SQLQueryError)
		}

		rows, err := res.RowsAffected()
		if err != nil {
			fmt.Printf("Failed to insert record. Error %s", err.Error())
			os.Exit(SQLQueryError)
		}

		fmt.Printf("Inserted %d row! %v\n", rows, v)
	}
}

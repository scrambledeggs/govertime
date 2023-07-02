package main

import (
	"database/sql"
	"flag"
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
)

func printTable(rows *sql.Rows) {
	columns, _ := rows.Columns()
	columnCount := len(columns)

	values := make([]string, columnCount)
	valuePtrs := make([]interface{}, columnCount)

	for i := range values {
		valuePtrs[i] = &values[i]
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(columns)

	for rows.Next() {
		err := rows.Scan(valuePtrs...)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		table.Append(values)
	}

	table.Render()
}

func handlerQueryError(err error, query string) {
	if err != nil {
		fmt.Printf("Encountered error while executing the query.\nQuery: %s\nError:%s\n", query, err.Error())
		os.Exit(SQLQueryError)
	}
}

func handleFlags() bool {
	db := openSql()
	defer db.Close()

	lsPtr := flag.Bool("ls", false, "List all overtime for the current month")
	gdtbPtr := flag.Bool("gdtb", false, "Include 29/30/31 from prev month")

	flag.Parse()

	if *lsPtr {
		if *gdtbPtr {
			rows, err := db.Query(ViewMonthGetDatThirtyBroOvertimeQuery)
			handlerQueryError(err, ViewMonthGetDatThirtyBroOvertimeQuery)
			defer rows.Close()
			printTable(rows)
			return true
		}

		rows, err := db.Query(ViewMonthOvertimeQuery)
		handlerQueryError(err, ViewMonthOvertimeQuery)
		defer rows.Close()
		printTable(rows)

		return true
	}

	return false
}

func handleInsert() {
	args := os.Args[1:]
	ot, err := govertime(args)
	if err != nil {
		fmt.Printf("encountered error: %s", err.Error())
		os.Exit(69)
	}

	db := openSql()
	defer db.Close()

	if _, err := db.Exec(CreateTableQuery); err != nil {
		handlerQueryError(err, CreateTableQuery)
		os.Exit(SQLQueryError)
	}

	for _, v := range ot {
		res, err := db.Exec(InsertOvertimeQuery, v.Name, v.TimeIn, v.TimeOut, v.HoursOT, v.Reason)
		handlerQueryError(err, InsertOvertimeQuery)

		rows, err := res.RowsAffected()
		if err != nil {
			fmt.Printf("Failed to insert record. Error %s", err.Error())
			os.Exit(SQLQueryError)
		}

		fmt.Printf("Inserted %d row! %v\n", rows, v)
	}
}

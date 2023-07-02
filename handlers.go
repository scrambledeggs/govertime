package main

import (
	"database/sql"
	"flag"
	"fmt"
	"os"
	"strings"

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

func handleQueryError(err error, query string) {
	if err != nil {
		fmt.Printf("Encountered error while executing the query.\nQuery: %s\nError:%s\n", query, err.Error())
		os.Exit(SQLQueryError)
	}
}

func handleQueryOutput(rows *sql.Rows, err error) {
	handleQueryError(err, ViewMonthGetDatThirtyBroOvertimeQuery)
	printTable(rows)
	rows.Close()
}

func handleFlags() bool {
	db := openSql()
	defer db.Close()

	lsPtr := flag.Bool("ls", false, "List all overtime for the current month")
	gdtbPtr := flag.Bool("gdtb", false, "Include 29/30/31 from prev month")
	namesPtr := flag.String("names", "all", "Filter list by name")

	flag.Parse()

	if *lsPtr {
		hasNames := *namesPtr != "all"

		if hasNames {
			strName := *namesPtr
			strName = strings.ReplaceAll(strName, " ", "")
			arrNames := strings.Split(strName, ",")

			sqlNames := make([]interface{}, 0, len(arrNames))
			placeholders := make([]string, 0, len(arrNames))
			for _, v := range arrNames {
				placeholders = append(placeholders, "?")
				sqlNames = append(sqlNames, v)
			}

			finalQuery := fmt.Sprintf(ViewMonthOvertimeWithNamesQueryFmt, strings.Join(placeholders, ","))
			if *gdtbPtr {
				finalQuery = fmt.Sprintf(ViewMonthGetDatThirtyBroOvertimeQueryFmt, strings.Join(placeholders, ","))
			}

			rows, err := db.Query(finalQuery, sqlNames...)
			handleQueryOutput(rows, err)

			return true
		}

		if *gdtbPtr {
			rows, err := db.Query(ViewMonthGetDatThirtyBroOvertimeQuery)
			handleQueryOutput(rows, err)
			return true
		}

		rows, err := db.Query(ViewMonthOvertimeQuery)
		handleQueryOutput(rows, err)

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
		handleQueryError(err, CreateTableQuery)
		os.Exit(SQLQueryError)
	}

	for _, v := range ot {
		res, err := db.Exec(InsertOvertimeQuery, v.Name, v.TimeIn, v.TimeOut, v.HoursOT, v.Reason)
		handleQueryError(err, InsertOvertimeQuery)

		rows, err := res.RowsAffected()
		if err != nil {
			fmt.Printf("Failed to insert record. Error %s", err.Error())
			os.Exit(SQLQueryError)
		}

		fmt.Printf("Inserted %d row! %v\n", rows, v)
	}
}

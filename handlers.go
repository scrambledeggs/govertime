package main

import (
	"flag"
	"fmt"
)

func handleFlags() bool {
	db := openSql()
	defer db.Close()

	lsPtr := flag.Bool("ls", false, "List all overtime for the current month")
	gdtbPtr := flag.Bool("gdtb", false, "Include 30/31 from prev month")

	flag.Parse()

	if *lsPtr {
		if *gdtbPtr {
			fmt.Printf("Current OTs for the month + 30/31\n")
			return true
		}

		// TODO: display table
		fmt.Printf("Current OTs for the month\n")

		return true
	}

	return false
}

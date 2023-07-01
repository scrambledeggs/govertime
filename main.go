package main

import (
	_ "modernc.org/sqlite"
)

func main() {
	shouldExit := handleFlags()
	if shouldExit {
		return
	}

	handleInsert()
}

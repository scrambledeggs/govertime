package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	ot, err := govertime(args)
	if err != nil {
		fmt.Printf("encountered error: %s", err.Error())
		os.Exit(69)
	}
	fmt.Printf("Overtime Obj %v\n", ot)
}

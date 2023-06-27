package main

import (
	"fmt"
	"os"
)

func govertime(args []string) string {
	return "Successfully recorded!"
}

func main() {
	args := os.Args[1:]
	fmt.Printf("Hello %s!\n", args[0])
}

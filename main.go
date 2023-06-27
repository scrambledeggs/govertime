package main

import (
	"errors"
	"fmt"
	"os"
)

func govertime(args []string) (string, error) {
	if len(args) != 6 {
		return "", errors.New("Incomplete arguments")
	}

	return "Successfully recorded!", nil
}

func main() {
	args := os.Args[1:]
	fmt.Printf("Hello %s!\n", args[0])
}

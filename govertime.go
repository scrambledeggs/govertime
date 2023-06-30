package main

import (
	"errors"
	"strings"
)

func govertime(args []string) ([]string, error) {
	if len(args) != 6 {
		return []string{}, errors.New("Incomplete arguments")
	}

	names := args[0]
	return strings.Split(names, ","), nil
}

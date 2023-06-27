package main

import "errors"

func govertime(args []string) (string, error) {
	if len(args) != 6 {
		return "", errors.New("Incomplete arguments")
	}

	return "Successfully recorded!", nil
}

package main

import "strings"

func parseNames(names string) []string {
	return strings.Split(names, ",")
}

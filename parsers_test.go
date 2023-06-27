package main

import (
	"testing"
)

func TestParseNames(t *testing.T) {
	t.Run("should be able to split names into array", func(t *testing.T) {
		names := "jett,robin,andres"
		want := []string{"jett", "robin", "andres"}

		got := parseNames(names)
		if len(got) != len(want) {
			t.Errorf("want %v, got %v", want, got)
		}

		for i, s := range got {
			if s != want[i] {
				t.Errorf("not parsed properly %v, %v", got, want)
			}
		}
	})
}

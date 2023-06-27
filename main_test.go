package main

import "testing"

func TestGovertime(t *testing.T) {
	t.Run("should be able to parse arguments", func(t *testing.T) {
		args := []string{
			"jett",
			"06-27-2023",
			"1700",
			"06-28-2023",
			"0100",
			"OP deployment",
		}

		want := "Successfully recorded!"
		resp, _ := govertime(args)
		if resp != want {
			t.Errorf("Expecting '%s', but got %s", want, resp)
		}
	})

	t.Run("should fail if arguments are lacking", func(t *testing.T) {
		args := []string{
			"hohoho",
			"look at me im lacking",
		}

		_, err := govertime(args)
		if err == nil {
			t.Errorf("must error with incomplete arguments")
		}
	})

	t.Run("should be able to accept multiple name entries", func(t *testing.T) {

	})
}

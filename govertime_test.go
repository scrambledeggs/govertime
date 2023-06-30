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

		_, err := govertime(args)
		if err != nil {
			t.Errorf("unable to parse arguments")
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

	t.Run("should return array of result based on first param", func(t *testing.T) {
		args := []string{
			"jett,maw,wayne",
			"06-27-2023",
			"1700",
			"06-28-2023",
			"0100",
			"OP deployment",
		}

		res, err := govertime(args)
		if err != nil {
			t.Errorf("encountered error while parsing arguments %s", err)
		}

		want := 3
		got := len(res)
		if got != want {
			t.Errorf("expecting %d count but got %d", want, got)
		}
	})
}

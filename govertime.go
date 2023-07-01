package main

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

type Overtime struct {
	Name    string
	TimeIn  time.Time
	TimeOut time.Time
	HoursOT float64
	Reason  string
}

func govertime(args []string) ([]Overtime, error) {
	if len(args) != 6 {
		return []Overtime{}, errors.New("Incomplete arguments")
	}

	names := args[0]
	names = strings.ReplaceAll(names, " ", "")
	entries := strings.Split(names, ",")

	overtimes := make([]Overtime, 0, len(entries))

	location := time.FixedZone("GMT+8", 8*60*60)
	layout := "01-02-2006 3:04PM"

	timeIn, err := time.ParseInLocation(layout, fmt.Sprintf("%s %s", args[1], args[2]), location)
	timeOut, err := time.ParseInLocation(layout, fmt.Sprintf("%s %s", args[3], args[4]), location)

	if err != nil {
		return overtimes, err
	}

	hoursOt := timeOut.Sub(timeIn).Hours()

	for _, v := range entries {
		overtimes = append(overtimes, Overtime{
			Name:    v,
			TimeIn:  timeIn,
			TimeOut: timeOut,
			HoursOT: hoursOt,
			Reason:  args[5],
		})
	}
	return overtimes, nil
}

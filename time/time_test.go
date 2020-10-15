package main

import (
	"fmt"
	"time"
)

func Example_time_now() {
	// func Now() Time
	fmt.Println(time.Now())
	// func Date(year int, month Month, day, hour, min, sec, nsec int, loc *Location) Time
	fmt.Println(time.Date(2018, 9, 12, 12, 30, 10, 0, time.Local))
	// func Parse(layout, value string) (Time, error)
	fmt.Println(time.Parse("2016-01-02 15:04:05", "2018-04-23 12:24:51"))
	// func ParseInLocation(layout, value string, loc *Location) (Time, error) (layout已带时区时可直接用Parse)
	fmt.Println(time.ParseInLocation("2006-01-02 15:04:05", "2017-05-11 14:06:06", time.Local))
	// Output: 2020
}

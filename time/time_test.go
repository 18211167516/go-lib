package main

import (
	"fmt"
	"time"
)

func Example_timeDate() {
	fmt.Printf("date = %v\n", time.Date(2018, 9, 12, 12, 30, 10, 0, time.Local))
	parse, _ := time.Parse("2006-01-02 15/04/05", "2018-04-23 12/24/51")
	fmt.Printf("parse = %v\n", parse)
	parsrLo, _ := time.ParseInLocation("2006-01-02 15:04:05", "2017-05-11 14:06:06", time.Local)
	fmt.Printf("parsrLo = %v\n", parsrLo)
	// Output:
	// date = 2018-09-12 12:30:10 +0800 CST
	// parse = 2018-04-23 12:24:51 +0000 UTC
	// parsrLo = 2017-05-11 14:06:06 +800 CST
}

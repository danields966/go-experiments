package main

/*
Given datetime in format 1986-04-16T05:20:00+06:00
Convert it to UnixDate: Wed Apr 16 05:20:00 +0600 1986
*/

import (
	"fmt"
	"time"
)

func main() {
	var task string
	if _, err := fmt.Scan(&task); err != nil {
		panic(err)
	}

	date, err := time.Parse(time.RFC3339, task)
	if err != nil {
		panic(err)
	}

	fmt.Print(date.Format(time.UnixDate))
}

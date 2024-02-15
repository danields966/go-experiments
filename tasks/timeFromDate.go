package main

/*
Given the duration in the format "12 min. 12 sec."
Add this duration to the given date (or current date)
*/

import (
	"fmt"
	"time"
)

func main() {
	var m, s int
	if _, err := fmt.Scanf("%d min. %d", &m, &s); err != nil {
		panic(err)
	}

	const now = 1589570165
	dt := time.Unix(now, 0) // or time.Now()
	diff, _ := time.ParseDuration(fmt.Sprintf("%dm%ds", m, s))
	fmt.Println(dt.Add(diff).Format(time.UnixDate))
}

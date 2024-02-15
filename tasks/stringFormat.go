package main

import "fmt"

/*
The hour hand has turned d degrees since the beginning of the day.
Determine how many whole hours h and whole minutes m there are now.
The program input is an integer d (0 < d < 360).
Print the text:
It is ... hours ... minutes.
*/

const (
	degInHour  = 30
	minInDeg   = 2
	timeFormat = "It is %d hours %d minutes."
)

func main() {
	var input int32
	if _, err := fmt.Scan(&input); err != nil {
		panic(err)
	}

	hours := input / degInHour
	minutes := (input % degInHour) * minInDeg

	fmt.Printf(timeFormat, hours, minutes)
}

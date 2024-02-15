package main

/*
Find the difference between the two given dates in format
13.03.2018 14:00:15,12.03.2018 14:00:15
*/

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	s := strings.Split(input, ",")
	t1, _ := time.Parse("02.01.2006 15:04:05", s[0])
	t2, _ := time.Parse("02.01.2006 15:04:05", s[1])

	var diff time.Duration
	if t2.After(t1) {
		diff = t2.Sub(t1)
	} else {
		diff = t1.Sub(t2)
	}
	fmt.Println(diff)
}

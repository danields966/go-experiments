package main

/*
Given the list of numbers in the range 0-100, each number is written on a new line (the total count is unknown).
You need to read all these numbers and print their sum to the Stdout.
*/

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
)

func createTestFile(filename string) {
	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := f.WriteString("1\n2\n3\n4\n5"); err != nil {
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	filename := "data/test.txt"
	createTestFile(filename)
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer func(f *os.File) {
		if err := f.Close(); err != nil {
			log.Fatal(err)
		}
	}(f)

	scanner := bufio.NewScanner(f)
	//scanner := bufio.NewScanner(os.Stdin)  // If we read from os.Stdin

	amount := 0
	for scanner.Scan() {
		a, _ := strconv.Atoi(scanner.Text())
		amount += a
	}
	if _, err := io.WriteString(os.Stdout, strconv.Itoa(amount)); err != nil {
		log.Fatal(err)
	}
}

package main

/*
Given the file with information about groups and students with their ratings.
You need to read the data and calculate the average number of ratings received by the students in the group.
The answer must be written to file or Stdout in JSON with the following format:
{
    "Average": 14.1
}
*/

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type (
	Group struct {
		Students []struct {
			Rating []int
		}
	}

	AverageRating struct {
		Average float64
	}
)

func main() {
	var group Group

	filename := "data/students.json"
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	inFile := f // inFile = os.Stdin
	decoder := json.NewDecoder(inFile)
	if err := decoder.Decode(&group); err != nil {
		panic(err)
	}

	var avg float64
	for _, st := range group.Students {
		avg += float64(len(st.Rating))
	}
	avg /= float64(len(group.Students))
	result := AverageRating{avg}

	outputFilename := "data/test.json"
	f1, err := os.Create(outputFilename)
	if err != nil {
		panic(err)
	}

	outFile := f1 // outFile := os.Stdout
	d, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		panic(err)
	}
	if _, err := io.WriteString(outFile, fmt.Sprintf("%s", d)); err != nil {
		panic(err)
	}

	if err := f.Close(); err != nil {
		panic(err)
	}
	if err := f1.Close(); err != nil {
		panic(err)
	}
}

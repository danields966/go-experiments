package main

import (
	"encoding/json"
	"fmt"
)

// In common struct tags looks like this: `json:"name,option,option"`

type playerStruct struct {
	// Use "name" instead of "Name" in marshal/unmarshal
	Name string `json:"name"`

	// Skip this field if it has an empty value (false, nil, an empty slice, etc.)
	Age int `json:",omitempty"`

	// Always skip this field in marshal/unmarshal
	Status bool `json:"-"`
}

func main() {
	m := playerStruct{Name: "John Connor", Age: 0, Status: true}

	data, err := json.Marshal(m)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s", data) // {"name":"John Connor"}
}

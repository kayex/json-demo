package main

import (
	"encoding/json"
	"io/ioutil"
)

type Person struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

func main() {
	// Read entire file contents to memory
	data, _ := ioutil.ReadFile("data.json")

	// Allocate Person
	p := Person{}

	// Decode json into Person
	err := json.Unmarshal(data, &p)
	if err != nil {
		panic(err)
	}

	// Update age
	p.Age = 42

	// Decode Person into json
	out, _ := json.Marshal(p)

	// Write new json to file
	err = ioutil.WriteFile("new-data.json", out, 0644)
	if err != nil {
		panic(err)
	}
}

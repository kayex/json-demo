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
	useMap()
}

func useStruct() {
	// Read entire file contents to memory.
	data, _ := ioutil.ReadFile("data.json")

	// Allocate Person.
	p := Person{}

	// Decode json into Person.
	err := json.Unmarshal(data, &p)
	if err != nil {
		panic(err)
	}

	// Update age.
	p.Age = p.Age + 5

	// Encode Person into json.
	out, _ := json.Marshal(p)

	// Write new json to file.
	err = ioutil.WriteFile("new-data.json", out, 0644)
	if err != nil {
		panic(err)
	}
}

func useMap() {
	// Read entire file contents to memory.
	data, _ := ioutil.ReadFile("data.json")

	// Create map.
	var p map[string]interface{}

	// Decode json.
	err := json.Unmarshal(data, &p)

	// Read age, and cast it from interface{} to float64.
	//
	// All json numbers have to be cast to float64 initially, as there are no
	// separate types for integers and floating point values in json.
	// Further conversion would have to be done with another step, but for demonstration purposes
	// we can let it remain as a float64.
	age := p["age"].(float64)

	// Update age
	p["age"] = age + 5.0

	// Encode map into json
	out, _ := json.Marshal(p)

	// Write new json to file.
	err = ioutil.WriteFile("new-data.json", out, 0644)
	if err != nil {
		panic(err)
	}
}

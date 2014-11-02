package parser

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

/**
 * Struct for JSON-scheme properties map
 */
type SchProperties map[string]Scheme

/**
 * Struct for holding a JSON-scheme
 */
type Scheme struct {
	Type       string        `json:type`
	Schema     string        `json:$schema`
	Id         string        `json:id`
	Required   bool          `json:required`
	Properties SchProperties `json:properties`
}

/**
 * Check for errors
 */
func check(e error) {
	if e != nil {
		panic(e)
	}
}

/**
 * Unmarshal JSON into &Scheme struct
 */
func Parse(a string) *Scheme {
	str := &Scheme{}
	json.Unmarshal([]byte(a), str)
	return str
}

/**
 * Read a JSON schema
 */
func ReadSchema(a string) string {
	dat, err := ioutil.ReadFile(a)
	check(err)
	//fmt.Print(string(dat))
	return string(dat)
}

func ParseSchema(a string) *Scheme {
	json_scheme := parser.ReadSchema(a)
	parsed_scheme := parser.Parse(json_scheme)
	return parsed_scheme
}

func main() {
	// Read a json schema
	json_scheme := ReadSchema("./schema/example1.json")

	// Parse the schema
	parsed_scheme := Parse(json_scheme)

	// Print the structure
	fmt.Printf("%+v", parsed_scheme)

}

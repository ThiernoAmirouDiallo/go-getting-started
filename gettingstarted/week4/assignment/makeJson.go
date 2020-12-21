package main

import (
	"bufio"
	json2 "encoding/json"
	"fmt"
	"os"
)

func main() {
	var name, address string

	askFor(&name, "name")
	askFor(&address, "address")

	m := make(map[string]string)

	m["name"] = name
	m["address"] = address

	json, _ := json2.Marshal(m)

	fmt.Printf("The json representation of %v is %v", m, string(json))

}

func askFor(s *string, field string) {
	for {
		scanner := bufio.NewScanner(os.Stdin)

		fmt.Printf("Please enter your %v : ", field)

		scanner.Scan()
		text := scanner.Text()

		if len(text) != 0 {
			*s = text
			break
		} else {
			fmt.Println("Error while reading your input")
		}
	}
}

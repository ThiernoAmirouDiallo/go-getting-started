package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

type Person struct {
	Fname string `json:"firstName"`
	Lname string `json:"lastName"`
}

func main() {
	var filename string
	var file *os.File
	var err error

	for {
		getUserInputFor(&filename, "file name")
		file, err = os.Open(filename)
		if err != nil {
			fmt.Printf("Error while opening the file %v : %v\n", filename, err)
		} else {
			break
		}
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var line []byte
	var persons []Person = make([]Person, 0, 3)

	for {
		line, _, err = reader.ReadLine()
		if err != nil {
			break
		}

		fAndLname := strings.Split(string(line), " ")
		person := Person{fAndLname[0], fAndLname[1]}

		persons = append(persons, person)
	}

	fmt.Printf("Printing the data read from the file %v\n", filename)
	for i, person := range persons {
		jsonData, err := json.Marshal(person)
		if err != nil {
			log.Println(err)
		}
		fmt.Printf("%v - %v\n", i, string(jsonData))
	}
}

func getUserInputFor(s *string, field string) {
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

package main

import (
	"fmt"
	"strings"
)

func main() {
	for {
		var enteredString string
		for {
			fmt.Print("Please enter a string : ")
			_, err := fmt.Scanf("%s", &enteredString)
			if err == nil {
				break
			}
		}

		if matchesSearchPattern(enteredString) {
			fmt.Println("Found!")
		} else {
			fmt.Println("Not Found!")

		}

		var contrinueResponse string
		fmt.Print("Do you want to enter another string Y/N ? ... ")
		fmt.Scan(&contrinueResponse)

		if !(strings.ToUpper(contrinueResponse) == "Y") {
			break
		}
	}

}

func matchesSearchPattern(s string) bool {
	lowerCaseS := strings.ToLower(s)
	return strings.HasPrefix(lowerCaseS, "i") &&
		strings.HasSuffix(lowerCaseS, "n") &&
		strings.Contains(lowerCaseS, "a")
}

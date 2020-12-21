package main

import (
	"fmt"
	"strings"
)

func main() {

	for {
		var decimalNum float32
		for {
			fmt.Print("Please enter a decimal number : ")
			_, err := fmt.Scan(&decimalNum)
			if err == nil {
				break
			}
			fmt.Println("You should enter a decimal number!")
		}
		fmt.Printf("You entered : %v and the truncated number is : %v\n", decimalNum, trunc(decimalNum))

		var response string
		fmt.Print("Do you want to enter another one Y/N ? ... ")
		fmt.Scan(&response)

		if !(strings.ToUpper(response) == "Y") {
			break
		}
	}

}

func trunc(decimalNumber float32) int {
	return int(decimalNumber)
}

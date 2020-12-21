package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	numbers := make([]int, 0, 3)
	for {
		var enteredString string
		for {
			fmt.Print("Please enter a number or X to finish : ")
			_, err := fmt.Scanf("%s", &enteredString)

			if err == nil {
				break
			} else {
				fmt.Println("Error while reading your input")
			}
		}

		if strings.ToUpper(enteredString) == "X" {
			fmt.Println(numbers)
			break
		}

		if n, errConv := strconv.Atoi(enteredString); errConv != nil {
			fmt.Println("Error while converting your input to integer")
			continue
		} else {
			numbers = append(numbers, n)
			sort.Ints(numbers)
			fmt.Println(numbers)
		}
	}
}

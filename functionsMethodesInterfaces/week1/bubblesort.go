package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	ints := GetPopulatedSlice()

	BubbleSort(ints)
	fmt.Println(ints)
}

func BubbleSort(ints []int) {
	for i := len(ints) - 1; i > 0; i-- {
		for j := 0; j <= i; j++ {
			if ints[j] > ints[i] {
				Swap(ints, i, j)
			}
		}
	}
}

func Swap(ints []int, i int, j int) {
	if i == j {
		return
	}

	ints[i] = ints[i] - ints[j]
	ints[j] = ints[i] + ints[j]
	ints[i] = ints[j] - ints[i]
}

func GetPopulatedSlice() []int {
	numbers := make([]int, 0, 10)
	for len(numbers) < 10 {
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
			break
		}

		if n, errConv := strconv.Atoi(enteredString); errConv != nil {
			fmt.Println("Error while converting your input to integer")
			continue
		} else {
			numbers = append(numbers, n)
		}
	}

	return numbers
}

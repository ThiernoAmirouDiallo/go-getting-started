package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	var acceleration, initialVelocity, initialDisplacement, time float64
	acceleration = askForFloat64("the acceleration")
	initialVelocity = askForFloat64("the initial velocity")
	initialDisplacement = askForFloat64("the initial displacement")
	time = askForFloat64("the time")

	fn := GenDisplaceFn(acceleration, initialVelocity, initialDisplacement)

	fmt.Printf("fn(%v) = %v\n", time, fn(time))
}

func GenDisplaceFn(acceleration, initialVelocity, initialDisplacement float64) func(float64) float64 {
	return func(t float64) float64 {
		return 0.5*acceleration*math.Pow(t, 2) + initialVelocity*t + initialDisplacement
	}
}

func askForFloat64(description string) float64 {
	for {
		scanner := bufio.NewScanner(os.Stdin)

		fmt.Printf("Please enter '%v' : ", description)

		scanner.Scan()
		text := scanner.Text()

		if len(text) != 0 {
			if decimalNum, errConv := strconv.ParseFloat(text, 64); errConv != nil {
				fmt.Println("Error while converting your input to integer")
			} else {
				return decimalNum
			}
		} else {
			fmt.Println("Error while reading your input")
		}
	}
}

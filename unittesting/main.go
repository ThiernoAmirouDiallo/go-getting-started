package main

import (
	"fmt"
	"github.com/ThiernoAmirouDiallo/go-getting-started/debounce"
)

var v int = 3

func main() {
	debounce.Example()
}

func main2() {
	fmt.Println("unit testing")
	fmt.Println(addTwo(3))

	fmt.Println("Pointer demo")
	fmt.Printf("file scope v = %v\n", v)

	pointerDemo()

	scanUserInput()
}

func addTwo(nb int) int {
	return nb + 2
}

func pointerDemo() {
	var i int = 1
	var p1 *int
	var p2 = new(int)
	var y int

	p1 = &i
	y = *p1
	*p2 = 4

	fmt.Printf("i = %v, &i = %v\n", i, &i)
	fmt.Printf("y = %v, &y = %v\n", y, &y)
	fmt.Printf("p1 = %v, *p1 = %v\n", p1, *p1)
	fmt.Printf("p2 = %v, *p2 = %v\n", p2, *p2)
}

func scanUserInput() {
	var appleNumber int

	fmt.Printf("What is the num of apple ? ")

	num, err := fmt.Scan(&appleNumber)

	if err != nil {
		fmt.Println("Error reading your input")
	}

	fmt.Printf("You entered %v words\n", num)
	fmt.Printf("The number you gave is %v\n", appleNumber)

}

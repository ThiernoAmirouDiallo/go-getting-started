package main

import (
	"fmt"
)

type Animal struct {
	name       string
	food       string
	locomotion string
	noise      string
}

func (animal *Animal) Eat() {
	fmt.Printf("The %v eats %v\n", animal.name, animal.food)
}

func (animal *Animal) Move() {
	fmt.Printf("The %v locomotion method is %v\n", animal.name, animal.locomotion)
}

func (animal *Animal) Speak() {
	fmt.Printf("The %v sound is %v\n", animal.name, animal.noise)
}

func main() {
	animals := map[string]Animal{
		"cow":   Animal{"cow", "grass", "walk", "moo"},
		"bird":  Animal{"bird", "worms", "fly", "peep"},
		"snake": Animal{"snake", "mice", "slither", "hsss"},
	}

	for {
		var animalName, request string
		fmt.Print("Please enter the animal name and the resquest. Example : cow eat > ")
		_, err := fmt.Scanf("%s %s", &animalName, &request)
		if err != nil {
			fmt.Println("Error while reading your input")
			continue
		}

		animal, ok := animals[animalName]

		if !ok {
			fmt.Printf("Can't find the animal named '%v'\n", animalName)
			continue
		}

		switch request {
		case "eat":
			animal.Eat()
		case "move":
			animal.Move()
		case "speak":
			animal.Speak()
		default:
			fmt.Printf("The request '%v' is invalid\n", request)
		}
	}
}

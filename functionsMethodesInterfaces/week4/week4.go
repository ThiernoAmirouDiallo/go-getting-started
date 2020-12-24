package main

import (
	"fmt"
)

type Request struct {
	name       string
	animalName string
	param      string
}

type AnimalInfo struct {
	food       string
	locomotion string
	noise      string
}

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
	name string
}

type Bird struct {
	name string
}

type Snake struct {
	name string
}

func (cow Cow) Eat() {
	fmt.Printf(descriptionFormarts["food"], "cow", cow.name, animalsInfo["cow"].food)
}

func (cow Cow) Move() {
	fmt.Printf(descriptionFormarts["locomotion"], "cow", cow.name, animalsInfo["cow"].locomotion)
}

func (cow Cow) Speak() {
	fmt.Printf(descriptionFormarts["noise"], "cow", cow.name, animalsInfo["cow"].noise)
}

func (bird Bird) Eat() {
	fmt.Printf(descriptionFormarts["food"], "bird", bird.name, animalsInfo["bird"].food)
}

func (bird Bird) Move() {
	fmt.Printf(descriptionFormarts["locomotion"], "bird", bird.name, animalsInfo["bird"].locomotion)
}

func (bird Bird) Speak() {
	fmt.Printf(descriptionFormarts["noise"], "bird", bird.name, animalsInfo["bird"].noise)
}

func (snake Snake) Eat() {
	fmt.Printf(descriptionFormarts["food"], "snake", snake.name, animalsInfo["snake"].food)
}

func (snake Snake) Move() {
	fmt.Printf(descriptionFormarts["locomotion"], "snake", snake.name, animalsInfo["snake"].locomotion)
}

func (snake Snake) Speak() {
	fmt.Printf(descriptionFormarts["noise"], "snake", snake.name, animalsInfo["snake"].noise)
}

var animalsInfo = map[string]AnimalInfo{
	"cow":   {"grass", "walk", "moo"},
	"bird":  {"worms", "fly", "peep"},
	"snake": {"mice", "slither", "hsss"},
}

var descriptionFormarts = map[string]string{
	"food":       "The %v %v eats %v\n",
	"locomotion": "The %v %v locomotion method is %v\n",
	"noise":      "The %v %v sound is %v\n",
}

var animals = make(map[string]Animal)

func main() {
	printHelp()

	for {
		processRequest(askAndGetRequest())
	}
}

func askAndGetRequest() Request {
	for {
		var request, animalName, param string

		fmt.Print("Please enter : 'request animalName param' > ")
		_, err := fmt.Scanf("%s %s %s", &request, &animalName, &param)
		if err != nil {
			fmt.Println("Error while reading your input")
			continue
		}

		return Request{request, animalName, param}
	}
}

func printHelp() {
	fmt.Println("Resquest examples :")
	fmt.Println("\t> newanimal spirit cow|bird|snake")
	fmt.Println("\t> query spirit eat|move|speak\n")
}

func processRequest(request Request) {
	switch request.name {
	case "newanimal":
		processNewAnimalRequest(request)
	case "query":
		processAnimalRequest(request)
	default:
		fmt.Printf("The request '%v' is invalid. The Authorized values are : newanimal and query\n", request.name)
	}
}

func processAnimalRequest(request Request) {
	if fetchedAnimal, ok := animals[request.animalName]; ok {
		showAnimalInfo(fetchedAnimal, request.param)
	} else {
		fmt.Printf("Can't find the animal named '%v'\n", request.animalName)
	}
}

func processNewAnimalRequest(request Request) {
	if newAnimal, err := newAnimal(request.animalName, request.param); err == nil {
		animals[request.animalName] = newAnimal
	} else {
		fmt.Printf(err.Error())
	}
}

func showAnimalInfo(animal Animal, requestParam string) {
	switch requestParam {
	case "eat":
		animal.Eat()
	case "move":
		animal.Move()
	case "speak":
		animal.Speak()
	default:
		fmt.Printf("The request '%v' is invalid. The authorized values are : eat, move and speak\n", requestParam)
	}
}

func newAnimal(animalName string, animalType string) (Animal, error) {
	switch animalType {
	case "cow":
		return Cow{animalName}, nil
	case "bird":
		return Bird{animalName}, nil
	case "snake":
		return Snake{animalName}, nil
	default:
		err := fmt.Errorf("The animal type '%v' is invalid. The Authorized values are : cow, bird and snake\n", animalType)
		return nil, err
	}
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
	food       string
	locomotion string
	noise      string
}

type Bird struct {
	food       string
	locomotion string
	noise      string
}

type Snake struct {
	food       string
	locomotion string
	noise      string
}

func (a Cow) Eat() {
	fmt.Println(a.food)
}

func (a Cow) Move() {
	fmt.Println(a.locomotion)
}

func (a Cow) Speak() {
	fmt.Println(a.noise)
}

func (a Bird) Eat() {
	fmt.Println(a.food)
}

func (a Bird) Move() {
	fmt.Println(a.locomotion)
}

func (a Bird) Speak() {
	fmt.Println(a.noise)
}

func (a Snake) Eat() {
	fmt.Println(a.food)
}

func (a Snake) Move() {
	fmt.Println(a.locomotion)
}

func (a Snake) Speak() {
	fmt.Println(a.noise)
}

type AnimalInfo struct {
	animalInfo map[string]Animal
}

func (a AnimalInfo) handleNewAnimal(animalName, animalType string) {

	switch animalType {
	case "cow":
		a.animalInfo[animalName] = Cow{"grass", "walk", "moo"}
		fmt.Println("Created it!")
	case "bird":
		a.animalInfo[animalName] = Bird{"worms", "fly", "peep"}
		fmt.Println("Created it!")
	case "snake":
		a.animalInfo[animalName] = Snake{"mice", "slither", "hsss"}
		fmt.Println("Created it!")
	default:
		fmt.Println("Invalid query. Try again..!")
	}
}

func (a AnimalInfo) handleQuery(animalName, query string) {

	v, exist := a.animalInfo[animalName]
	if !exist {
		fmt.Println("Animal doesn't exist")
	}
	switch query {
	case "eat":
		v.Eat()
	case "move":
		v.Move()
	case "speak":
		v.Speak()
	default:
		fmt.Println("Invalid query. Try again..!")
	}
}

func main() {

	animalInfo := AnimalInfo{make(map[string]Animal)}
	for {
		fmt.Println("\nCreate a new animal or query \n1. newanimal <animal_name> <animal_type> : Ex newanimal joe cow\n1. query <animal_name> <query_type> : Ex query joe speak\n (x for exit)")
		fmt.Print(">")

		in := bufio.NewReader(os.Stdin)

		input, err := in.ReadString('\n')
		if err != nil {
			log.Printf("Error reading input, %v", err)
			return
		}

		args := strings.Split(input, " ")
		commandType := strings.TrimSpace(args[0])

		switch commandType {
		case "newanimal":
			animalInfo.handleNewAnimal(strings.TrimSpace(args[1]), strings.TrimSpace(args[2]))
		case "query":
			animalInfo.handleQuery(strings.TrimSpace(args[1]), strings.TrimSpace(args[2]))
		case "x":
			fmt.Println("Exiting..!")
			return
		default:
			fmt.Println("Invalid query. Try again..!")
		}

	}
}

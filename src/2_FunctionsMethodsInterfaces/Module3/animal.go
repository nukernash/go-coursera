package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (a Animal) Eat() {
	fmt.Println(a.food)
}

func (a Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a Animal) Speak() {
	fmt.Println(a.noise)
}

func initialize() map[string]Animal {
	return map[string]Animal{
		"cow":   {"grass", "walk", "moo"},
		"bird":  {"worms", "fly", "peep"},
		"snake": {"mice", "slither", "hsss"},
	}
}

func main() {

	animalMap := initialize()
	for {
		fmt.Println("Enter a animal name(cow, bird or snake) and a query(eat, move or speak) space seperated. (x for exit)")

		fmt.Print(">")

		in := bufio.NewReader(os.Stdin)

		input, err := in.ReadString('\n')
		if err != nil {
			log.Printf("Error reading input, %v", err)
			return
		}

		if strings.TrimSpace(input) == "x" {
			fmt.Println("Exiting..!")
			return
		}
		args := strings.Split(input, " ")
		animal := strings.TrimSpace(args[0])
		query := strings.TrimSpace(args[1])

		_, exist := animalMap[animal]
		if !exist {
			fmt.Println("Invalid animal")
			continue
		}
		switch query {
		case "eat":
			animalMap[animal].Eat()
		case "move":
			animalMap[animal].Move()
		case "speak":
			animalMap[animal].Speak()
		default:
			fmt.Println("Invalid query. Try again..!")
		}
	}
}

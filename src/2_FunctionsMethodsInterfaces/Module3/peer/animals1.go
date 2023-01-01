package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

var ANIMALS = map[string]Animal{
	"cow":   {food: "grass", locomotion: "walk", noise: "moo"},
	"bird":  {food: "worms", locomotion: "fly", noise: "peep"},
	"snake": {food: "mice", locomotion: "slither", noise: "hsss"},
}

type UserRequest struct {
	animalName   string
	animalAction string
}

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

func ParseUserRequest(input string) (*UserRequest, error) {
	fields := strings.Fields(input)
	if len(fields) != 2 {
		return nil, errors.New("Request should contain 2 words separated by space")
	}
	return &UserRequest{
		animalName:   strings.TrimSpace(fields[0]),
		animalAction: strings.TrimSpace(fields[1]),
	}, nil
}

func ExecuteUserRequest(request *UserRequest) error {
	animal, exists := ANIMALS[request.animalName]
	if !exists {
		return errors.New(fmt.Sprintf("Animal '%s' doesn't exist!", request.animalName))
	}
	switch request.animalAction {
	case "eat":
		animal.Eat()
	case "move":
		animal.Move()
	case "speak":
		animal.Speak()
	default:
		return errors.New(fmt.Sprintf("Action '%s' doesn't exist!", request.animalAction))
	}
	return nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter animal name and action separated by space or X to exit")
	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading user input")
			continue
		}
		if strings.TrimSpace(input) == "X" {
			os.Exit(0)
		}
		request, err := ParseUserRequest(input)
		if err != nil {
			fmt.Println("Error parsing user input")
			continue
		}
		err = ExecuteUserRequest(request)
		if err != nil {
			fmt.Println(fmt.Errorf("Error executing action: %w", err))
			continue
		}
	}
}

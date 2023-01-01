package main

import (
	"fmt"
	"strings"
)

//Animal interface
type Animal interface {
	Eat()
	Move()
	Speak()
}

//Cow of animal struct
type Cow struct {
}

//Eat cow
func (a Cow) Eat() {
	fmt.Println("grass")
}

//Move cow
func (a Cow) Move() {
	fmt.Println("walk")
}

//Speak cow
func (a Cow) Speak() {
	fmt.Println("moo")
}

//Bird b
type Bird struct {
}

//Eat bird
func (a Bird) Eat() {
	fmt.Println("worms")
}

//Move Bird
func (a Bird) Move() {
	fmt.Println("fly")
}

//Speak Bird
func (a Bird) Speak() {
	fmt.Println("peep")
}

//Snake s
type Snake struct {
}

//Eat Snake
func (a Snake) Eat() {
	fmt.Println("mice")
}

//Move Snake
func (a Snake) Move() {
	fmt.Println("slither")
}

//Speak Snake
func (a Snake) Speak() {
	fmt.Println("hss")
}

//AnimalMap to hold unique animals
type AnimalMap map[string]Animal

func main() {
	animalMap := make(AnimalMap)
	for {
		fmt.Print(">")

		var command, nameAni, info string
		_, err := fmt.Scanf("%s %s %s", &command, &nameAni, &info)
		if err != nil {
			fmt.Println("Error, input should be in this format: newanimal/query name type/info", err.Error())
			continue
		}
		command = strings.ToUpper(command)
		nameAni = strings.ToUpper(nameAni)
		info = strings.ToUpper(info)
		switch command {
		case "NEWANIMAL":
			{
				created := true
				switch info {
				case "COW":
					animalMap[nameAni] = Cow{}
				case "BIRD":
					animalMap[nameAni] = Bird{}
				case "SNAKE":
					animalMap[nameAni] = Snake{}
				default:
					fmt.Println("Unknown animal type please try again")
					created = false
				}
				if created {
					fmt.Println("Created it!")
				}

			}
		case "QUERY":
			{
				animal, ok := animalMap[nameAni]

				if ok {
					switch info {
					case "EAT":
						animal.Eat()
					case "MOVE":
						animal.Move()
					case "SPEAK":
						animal.Speak()
					default:
						fmt.Println("Unknown information about animal")
					}

				} else {
					fmt.Println("Animal not found")
				}
			}
		}
	}

}

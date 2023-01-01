package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	eat   string
	move  string
	speak string
}

func (a Animal) Eat() {
	fmt.Println(a.eat)
}

func (a Animal) Move() {
	fmt.Println(a.move)
}

func (a Animal) Speak() {
	fmt.Println(a.speak)
}

func main() {
	var cow, bird, snake Animal
	cow.eat = "grass"
	bird.eat = "worms"
	snake.eat = "mice"
	cow.move = "walk"
	bird.move = "fly"
	snake.move = "slither"
	cow.speak = "moo"
	bird.speak = "peep"
	snake.speak = "hsss"
	reader := bufio.NewReader(os.Stdin)

	for true {
		fmt.Print(">")
		str, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error in input")
			os.Exit(0)
		}
		words := strings.Fields(str)
		if len(words) != 2 {
			fmt.Println("Error in input")
			os.Exit(0)
		}
		if strings.Compare(words[0], "cow") == 0 {
			checkSecondWord(words[1], cow)
		} else if strings.Compare(words[0], "bird") == 0 {
			checkSecondWord(words[1], bird)
		} else if strings.Compare(words[0], "snake") == 0 {
			checkSecondWord(words[1], snake)
		} else {
			fmt.Println("Error in input")
			os.Exit(0)
		}
	}
}

func checkSecondWord(word string, a Animal) {
	if strings.Compare(word, "eat") == 0 {
		a.Eat()
	} else if strings.Compare(word, "move") == 0 {
		a.Move()
	} else if strings.Compare(word, "speak") == 0 {
		a.Speak()
	} else {
		fmt.Println("Error in input")
		os.Exit(0)
	}

}

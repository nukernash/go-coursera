package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Print("Enter something: ")
	in := bufio.NewReader(os.Stdin)
	userInput, err := in.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	userInput = strings.ToLower(userInput)
	if userInput[0] == 'i' && userInput[len(userInput)-2] == 'n' && strings.Contains(userInput, "a") {
		fmt.Println("Found!")

	} else {
		fmt.Println("Not Found!")
	}

}

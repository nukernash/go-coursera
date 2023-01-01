package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	var input string
	fmt.Println("Enter String to find ian : ")

	in := bufio.NewReader(os.Stdin)

	input, err := in.ReadString('\n')
	input = strings.TrimSpace(input)
	fmt.Println("Input : " + input)

	if err != nil {
		fmt.Printf("Invalid Input: %s\n", err)
	}

	if (strings.ContainsAny(input, "a") || strings.ContainsAny(input, "A")) &&
		(strings.HasPrefix(input, "i") || strings.HasPrefix(input, "I")) &&
		(strings.HasSuffix(input, "n") || strings.HasSuffix(input, "N")) {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}

}

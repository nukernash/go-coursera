/*
Write a program which prompts the user to enter a string.
The program searches through the entered string for the characters ‘i’, ‘a’, and ‘n’.
The program should print “Found!” if the entered string starts with the character ‘i’,
ends with the character ‘n’, and contains the character ‘a’.
The program should print “Not Found!” otherwise.
The program should not be case-sensitive, so it does not matter if the characters are upper-case or lower-case.
*/
package main

import (
	"fmt"
	"strings"
)

func checkForIAN(s string) bool {
	sUppercase := strings.ToUpper(s)
	if strings.HasPrefix(sUppercase, "I") &&
		strings.HasSuffix(sUppercase, "N") &&
		strings.Contains(sUppercase, "A") {
		return true
	}
	return false
}

func main() {
	var input string
	fmt.Print("(ian) Enter a string: ")
	_, err := fmt.Scanf("%s", &input)
	if err != nil {
		fmt.Println("error reading string value", err)
	} else {
		if checkForIAN(input) {
			fmt.Println("Found!")
		} else {
			fmt.Println("Not Found!")
		}

	}
}

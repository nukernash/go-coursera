package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Name struct {
	fname string
	lname string
}

func main() {

	names := []Name{}
	fmt.Println("Enter a file name with names")

	var fileName string
	fmt.Scan(&fileName)

	file, err := os.Open(fileName)

	if err != nil {
		fmt.Println("error while reading file, $v", err)
		return
	}

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		fullName := strings.Split(scanner.Text(), " ")
		names = append(names, Name{fullName[0], fullName[1]})
	}

	for _, v := range names {
		fmt.Println(v.fname + " " + v.lname)
	}

}

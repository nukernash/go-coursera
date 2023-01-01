package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type person struct {
	fname string
	lname string
}

func main() {
	var fpath string
	// Get fpath input
	fmt.Print("Enter the name of the text file: ")
	_, err := fmt.Scan(&fpath)
	if err != nil {
		log.Fatal(err)
		return
	}
	// Read the file
	file, err := os.Open(fpath)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()
	// Scan file line by line
	people := make([]person, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		name := strings.Split(scanner.Text(), " ")
		if len(name) == 2 {
			// Limit size of name
			if len(name[0]) > 20 {
				name[0] = name[0][:20]
			}
			if len(name[1]) > 20 {
				name[1] = name[1][:20]
			}
			people = append(people, person{name[0], name[1]})
		}
	}
	// Print the people
	for _, person := range people {
		fmt.Println(person.fname, person.lname)
	}

}

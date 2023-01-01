package main

import (
	"fmt"
	"os"
	"strings"
)

type Person struct {
	fname string
	lname string
}

func main() {
	fmt.Println("please give the name of the txt file")
	var name string
	fmt.Scan(&name)
	read, err := os.ReadFile(name)
	if err == nil {
		var p1 Person
		sli := make([]Person, 0)
		lines := strings.Split(string(read), "\n")
		for _, line := range lines {
			name := strings.Split(line, " ")
			p1.fname = name[0]
			p1.lname = name[1]
			sli = append(sli, p1)
		}

		for i, v := range sli {
			fmt.Printf("%d)first name: %s , last name: %s \n", i+1, v.fname, v.lname)
		}
	} else {
		fmt.Println(err)
	}
}

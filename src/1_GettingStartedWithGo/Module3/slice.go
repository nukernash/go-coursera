package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {

	var arg string

	slice := make([]int, 0, 4)

	for {
		fmt.Println("Enter any Integer('x' to Exit) : ")
		_, err := fmt.Scan(&arg)

		if err != nil {
			fmt.Printf("Invalid Input: %s\n", err)
			continue
		}

		if arg == "x" {
			fmt.Println("Exiting.. !")
			break
		}

		i, err := strconv.Atoi(arg)
		if err != nil {
			fmt.Println("Enter only integer")
			continue
		}

		slice = append(slice, i)
		sort.Ints(slice)

		fmt.Printf("Sorted slice : %v\n", slice)
	}
}

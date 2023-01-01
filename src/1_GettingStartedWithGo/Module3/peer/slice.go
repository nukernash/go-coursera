package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	sli := make([]int, 3)
	var userInput string
	for {
		fmt.Print("Enter a Integer: ")
		// Read user input
		_, err := fmt.Scan(&userInput)
		if err != nil {
			fmt.Println(err)
		}
		// Break if input == X
		if userInput == "X" {
			break
		}
		// Convert to int
		num, err := strconv.Atoi(userInput)
		if err != nil {
			fmt.Println(err)
			continue
		}
		sli = append(sli, num)
		// Sort the slice
		sort.Sort(sort.IntSlice(sli))
		// Print slice
		fmt.Printf("%v\n", sli)
	}
}

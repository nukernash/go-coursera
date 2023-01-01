package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	slice := make([]int, 0, 3)
	for {
		var input string
		fmt.Print("Please enter a number")
		fmt.Scan(&input)
		i, err := strconv.Atoi(input)
		if err != nil {
			switch input {
			case "X":
				os.Exit(1)
			case "x":
				os.Exit(1)
			default:
				continue
			}
		}
		slice = append(slice, i)
		sort.Ints(slice)
		fmt.Println(slice)
	}
}

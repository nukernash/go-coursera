package main

import (
	"fmt"
	"strconv"
)

func main() {

	var input string
	for i := 0; i < 2; i++ {
		fmt.Println("Enter floating number to truncate : ")

		_, err := fmt.Scan(&input)

		f, err := strconv.ParseFloat(input, 2)
		if err != nil {
			fmt.Printf("Invalid Input: %s\n", err)
		}

		truncated := int64(f)
		fmt.Printf("Truncated integer value for %f : %d\n", f, truncated)

	}
}

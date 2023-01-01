package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func BubbleSort(input []int) {
	for i := len(input) - 1; i > 0; i-- {
		for j := 1; j <= i; j++ {
			if input[j] < input[j-1] {
				swap(input, j)
			}
		}
	}
}

func swap(input []int, x int) {
	temp := input[x]
	input[x] = input[x-1]
	input[x-1] = temp
}

func main() {

	fmt.Println("Enter Integers seperated by space to be sorted and hit enter")

	var input []int

	var arg string
	in := bufio.NewReader(os.Stdin)

	arg, err := in.ReadString('\n')
	if err != nil {
		log.Printf("Error reading input, %v", err)
		return
	}

	for _, v := range strings.Split(strings.TrimSpace(arg), " ") {
		i, err := strconv.Atoi(v)
		if err != nil {
			log.Printf("Invalid input, %v", err)
			return
		}
		input = append(input, i)
	}
	BubbleSort(input)
	fmt.Printf("Sorted array is -> %v\n", input)

}

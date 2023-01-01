package main

import (
	"fmt"
)

// Constants
const INPUT_LENGTH int = 10

func main() {
	inputDataArr := make([]int, INPUT_LENGTH)
	for i := 0; i < INPUT_LENGTH; i++ {
		fmt.Print("Enter an integer (", i+1, "): ")
		fmt.Scan(&inputDataArr[i])
	}

	fmt.Println("BEFORE SORTING : ", inputDataArr)

	BubbleSort(inputDataArr)

	fmt.Println("SORTED ARRAY IS:  ", inputDataArr)
}

func BubbleSort(inputArr []int) {
	for i := 0; i < len(inputArr)-1; i++ {
		for j := 0; j < len(inputArr)-1-i; j++ {
			Swap(inputArr, j)
		}
	}
}

func Swap(slice []int, index int) {
	if slice[index] > slice[index+1] {
		temp := slice[index+1]
		slice[index+1] = slice[index]
		slice[index] = temp
	}
}

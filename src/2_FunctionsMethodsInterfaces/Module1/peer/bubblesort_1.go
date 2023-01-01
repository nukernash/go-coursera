package main

import (
	"fmt"
	"strconv"
)

//BubbleSort ...
func BubbleSort(slice []int) {
	length := len(slice)
	for i := 0; i < length-1; i++ {
		for j := 0; j < length-i-1; j++ {
			if slice[j] > slice[j+1] {
				Swap(slice, j)
			}
		}
	}
}

//Swap ...
func Swap(slice []int, i int) {
	slice[i], slice[i+1] = slice[i+1], slice[i]
}

func main() {
	var slice = make([]int, 0, 10)
	var s string
	for i := 0; i < 10; i++ {
		fmt.Scan(&s)

		num, _ := strconv.Atoi(s)
		slice = append(slice, num)

		BubbleSort(slice)
		fmt.Println(slice)
	}

}

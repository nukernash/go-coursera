package main

import (
	"fmt"
	"math"
)

func GenDisplaceFn(acceleration, initial_velocity, initial_displacement float64) func(float64) float64 {

	fn := func(time float64) float64 {
		return (1 / 2 * acceleration * math.Pow(time, 2)) + (initial_velocity * time) + initial_displacement
	}
	return fn
}

func getInput(param string) float64 {
	fmt.Printf("Enter %s :\n", param)

	var input float64

	fmt.Scanf("%f", &input)

	return input
}

func main() {

	acceleration := getInput("acceleration")
	initial_velocity := getInput("initial_velocity")
	initial_displacement := getInput("initial_displacement")

	fn := GenDisplaceFn(acceleration, initial_velocity, initial_displacement)

	time := getInput("time")

	fmt.Printf("displacement : %f\n", fn(time))

}

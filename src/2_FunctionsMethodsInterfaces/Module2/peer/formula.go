package main

import "fmt"

func GenDisplaceFn(a, v0, s0 float64) func(t float64) float64 {
	fn := func(t float64) float64 {
		s := 0.5*a*t*t + v0*t + s0
		return s
	}
	return fn
}

func main() {
	var a, v0, s0, t float64
	fmt.Print("Enter acceleration: ")
	fmt.Scan(&a)
	fmt.Print("Enter initial velocity: ")
	fmt.Scan(&v0)
	fmt.Print("Enter initial displacement: ")
	fmt.Scan(&s0)

	fn := GenDisplaceFn(a, v0, s0)

	for {
		fmt.Print("Enter time: ")
		fmt.Scan(&t)
		fmt.Println("The displacement after the entered time is", fn(t))
	}
}

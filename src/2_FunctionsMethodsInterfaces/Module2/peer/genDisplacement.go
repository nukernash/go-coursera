package main

import (
	"fmt"
)

func inputFloat(valName string) float64 {
	var val float64
	fmt.Printf("Enter the value of %s: ", valName)
	_, err := fmt.Scan(&val)
	if err != nil{
		fmt.Printf("ERROR [%v] | Invalid Value \n", err) 
		return inputFloat(valName)
	}
	return val
}

func main() {

	a := inputFloat("Acceleration")
	v := inputFloat("Initial Velocity")
	s := inputFloat("Initial Displacement")	
	
	fn := GenDisplaceFn(a, v, s)
	t := inputFloat("Time")	

	fmt.Printf("Displacement after [%v] seconds => %v \n", t, fn(t))
}

func GenDisplaceFn(a, v0, s0 float64) func(float64) float64 {
	return func(t float64) float64 {
		return 0.5*a*t*t + v0*t + s0
	}
}
package main

import (
	"fmt"
	"math"
)

func main() {
	weightLbs := 614.0
	weightKg := weightLbs/2.2
	
	mass := weightKg/9.8
	
	//fmt.Println(mass)
	remainder := math.Remainder(mass, 1)
	fmt.Println(remainder)
	
	if remainder > 0.4 {
		mass = mass - remainder
		mass = mass + 1
	}
	
	
	
	fmt.Println(mass)
}

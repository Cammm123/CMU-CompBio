package main

import (
	"fmt"
)

func main() {
	fmt.Println("While Loops!")
	fmt.Println(Factorial(5))
	fmt.Println(Factorial(70))
	fmt.Println(SumFirstNIntegers(6))
}

//Factorial takes as input an integer n and returns n!
func Factorial(n int) int {
	if n < 0 {
		panic("Error: negative input given to Factorial().")
	}

	p := 1 //will store the producrt and severs as a counter
	i := 1

	for i <= n { //Go has no word for while, instead only has for
		p *= i
		i++
	}
	
	return p
}

/*
Challenge: Write a function taht takes input as an integer (SumFirstNIntegers)
	This should reutrn the sum of the frist n positive integers
*/

func SumFirstNIntegers(n int) int {

	if n < 0 {
		panic("Error: negative input given to SumFirstNIntegers")
	}

	sum := 0 //storing sum
	i := 1 //counter

	for i <= n {
		sum += i
		i++
	}

	return sum
}
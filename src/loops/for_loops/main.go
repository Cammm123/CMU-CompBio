package main

import (
	"fmt"
)

func main() {
	fmt.Println("For Loops!")

	fmt.Println(Factorial(4))
	fmt.Println(SumFirstNIntegers(100))
	fmt.Println(GaussSum(100))

	fmt.Println(SumEven(199))
	/*
	This is an example of underflow
	var count uint = 10
	for ; count >= 0; count-- {
		fmt.Println(count)
	}

	fmt.Println("Blast Off!")
	*/
}

//Factorial takes input as integer n and return n!
func Factorial(n int) int {

	if n < 0 {
		panic("Error: You cannot input a negative integer inside Factorial().")
	}

	p := 1
	
	for i := 1; i <=n; i++ {
		// i := 1 is called the "initialization step"
		// i <= n is called the "condition"
		// i++ is called the "post condition"

		p *= i
	}

	return p

}

func SumFirstNIntegers(n int) int {
	if n < 0 {
		panic("Error: You cannot input a negative integer inside SumFirstNIntegers().")
	}

	sum := 0

	for i := 1; i <= n; i++ {
		sum += i
	}

	return sum
}

func GaussSum(n int) int {
	return n*(n+1)/2
}

func YetAnotherFactorial(n int) int {
	if n < 0 {
		panic("Error: You cannot input a negative integer inside YetAnotherFactorial().")
	}
	
	p := 1
	
	for i := n; i>= n; i-- {
		p*=i
	}
	return p
}

//SumEven takes as input an integer k and returns the sum of all even positive integers up to k
func SumEven(n int) int {
	if n < 0 {
		panic("Error: You cannot input a negative integer into SumEven()")
	}

	sum := 0
	for i:=n; i >=n; i-- {
		if i % 2 == 0 { //checking to see that it is even
			sum += i
			i--
		} else {
			i--
		}
		
	}
	return sum
}
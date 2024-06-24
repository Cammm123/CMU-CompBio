package main

import (
	"fmt"
)

func main() {
	fmt.Println("Functions!")

	x := 3
	n := SumTwoInts(x, x)
	fmt.Println(n)

	m := 17
	fmt.Println(AddOne(m))
	fmt.Println(m) //what gets printed?
}

// AddOne takes an integer k as input and returns the value of k+1
func AddOne(k int) int {
	k = k+1
	return k
}

// Take two ints as inputs, sums them up, and spits out the sum
func SumTwoInts(a int, b int) int {
	return a + b
}

//Takes two input integers, and returns two copies of that inpur variable
func DoubleAndDuplicate(x float64) (float64, float64) {
	return 2.0 * x, 2.0*x
}

//Function Pi takes no inputs and returns the value of pi.
func Pi() float64 {
	return 3.14
}

//Fucntion PrintHi will print "Hi"
func PrintHi() {
	fmt.Println("Hi")
}
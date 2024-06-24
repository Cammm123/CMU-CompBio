package main
import (
	"fmt"
)

func main() {
	fmt.Println("Arrays (and slices).")

	var list [6]int

	list[0] = -8

	i:=3
	list[2*i-4] = 17

	//These commands lead to compiler errors
	//list[len(list)] = 7 --> out of bounds
	//list[-4] = 2 --> no negative index allowed in go

	list[len(list) - 1] = 43
	fmt.Println(list)

	//slice declarations are a little different
	var a []int //right now, a has a special value called "nil"
	n:=4

	//slices must be made
	a = make([]int, n)

	//we set value of arrays similar to those of slices
	a[0] = 6
	a[2] = -33

	//a one line delcaration
	b := make([]int, n+2)
	fmt.Println(b)

	fmt.Println(a)

	fmt.Println(FactorialArray(6))

	var c [6]int 
	d := make([]int, 6)

	ChangeFirstElementArray(c)
	ChangeFirstElementSlice(d)
	fmt.Println("c is now", c)
	fmt.Println("d is now", d)

	fmt.Println(MinIntegers(3, 6, 7, 1, 5))
}

//variadic functions take a cariable number of inputs

//MinIntegers takes as input an arbitrary number of integers and returns thei minimum value.
func MinIntegers(numbers ...int) int {
	if len(numbers) == 0 {
		panic("Error: empty slice given to MinIntegerArray().")
	}
	
	return MinIntegerArray(numbers)
}


//MinintegerArray takes as input a slice of integers and returns the minimum value in that slice
func MinIntegerArray(list []int) int {
	if len(list) == 0 {
		panic("Error: empty slice given to MinIntegerArray().")
	}
	
	min := list[0] //will store minum value

	for i:=0;i<len(list);i++ { //This is equivalent to "for i:=range list"
		if list[i] < min {
			min = list[i]
		}
	}

	return min
}

func ChangeFirstElementArray(a [6]int) {
	a[0] = 1 // copy of input array gets edited
	//copy is destroyed
}

func ChangeFirstElementSlice(a []int) {
	a[0] = 1
	//when you pass in a slice to a function, you get access to that underlying array
}

//FactorialArray takes as input an integer n, and it returns an array of length n+1 whose k-th elemnt is k!
func FactorialArray(n int) []int {
	if n < 0 {
		panic("Error: Negative input given to FactorialArray().")
	}

	factorial := make([]int, n+1)
 	//recall that 0! is 1
	factorial[0] =1

	//range k from 1 to n, and use the fact that k! = k*(k-1)
	for k:=1;k<=n;k++ {
		factorial[k] = k * factorial[k-1]
	}

	return factorial
}

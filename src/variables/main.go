package main

import "fmt"

// Go won't read this line

/*
this is a multi line long comment
*/

func main() {
	fmt.Println("Variables and types.")
	var integer_variable int = 14 // default value of interger is 0
	var float_variable float64 = -2.3 // default value of float is 0.0
	var string_variable string = "hi" // default value of string is ""
	var uint_variable uint = 14 // default value for uint is 0
	var byte_variable byte = 'H' // default value of byte is 0...SYMBOL IN SINGLE QUOTES
	var bool_variable bool = true // default value for boolean is "false"

	// shorthad declarations avoid lengthy var integer_variable int statements
	i := -6 // equivalent to var i int = -6, and i's type (int) is inferred
	hi := "Yo" // has type string
	k := 34 // automatically have type int
	y := 3.7 // has type float64
	secondStatement := true // has type bool

	// Lets see what values these variables have
	fmt.Println(integer_variable)
	fmt.Println(float_variable)
	fmt.Println(string_variable)
	fmt.Println(uint_variable)
	fmt.Println(byte_variable)
	fmt.Println(bool_variable)

	fmt.Println(i, hi, k, y, secondStatement)

	// Now going to try some math
	fmt.Println(2 * (i+5) * k)
	fmt.Println(2*y - 3.16)

	fmt.Println(float64(k) * y) // Casting it as float in order to do multiplication
	fmt.Println(k / 14) // expecting 2.4285...but we see 2!

	/* 
	Dividing two integers corresponds to integer division (throw away the remainder)
	If I want to get the remainder, I need to cast it as a float64
	*/

	fmt.Println(float64(k) / 14.0)

	// Not ever type of conversion will work...bool(0)

	var p int = -187
	var s uint = uint(p)
	fmt.Println(s)

	m := 9223372036854775807 // Probably see this 7 be replaced with an 8
	fmt.Println(m+1)

}


package main

import (
	"fmt"
)

func main() {
	fmt.Println("Conditionals!")
	fmt.Println("The minimum of 3 and 4 is ", Min2(3,4))
	fmt.Println(WhichIsGreater(3,5)) //Should be -1
	fmt.Println(WhichIsGreater(42,42)) //Should be 0
	fmt.Println(WhichIsGreater(-2,-7)) //Should be 1
}

//Min2 fucntion takes two integers as inputs and returns their minium.
func Min2(a, b int) int {
	if a < b{
		return a
	} else {
		// b must be smaller or equal, so we can return b
		return b
	}
}

//Positive Difference takes the input of two integers. It returns the absolute value
func PositiveDifference(a, b int) int {
	var c int
	if a==b {
		return 0
	} else if a > b {
		c= a - b
	} else {
		c = b - a //can saefly assume that b is more than a
	}
	return c
}


/*SameSign takes as input as two integers.
	True: if they are same sign
	False: if they have diff sign
*/
func SameSign(x, y int) bool {
	if x * y > 0 { //x*y must be positive if they are the same
		return true
	} else {
		return false
	}
}


/* WhichIsGreater compares two input integers and returns:
	1: if the first is larger
	-1: if the second is larger
	0: if they are equal
*/
func WhichIsGreater(a, b int) int {
	//we need three cases
	if a == b{
		return 0
	} else if a > b {
			return 1
	} else {
			//if we got here, we know that y>x
			return -1
	}

}


//Index of comparison operators
/*
> : more than
< : less than
>= : greater than or equal to
<= : less than or equal to
== : equal to
!= : not equal to
! : "not" For example, if x is boolean, then !x is false
*/
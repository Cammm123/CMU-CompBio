package main

import (
	"fmt"
	//"math"
)

func main() {
	fmt.Println(Permutation(12, 1))
	fmt.Println(Permutation(10, 10))
	fmt.Println(Combination(10, 4))
	fmt.Println(Combination(1000, 998))
}

func Factorial(n int) int {

	product := 1

	for p:=1; p<=n; p++ {
		product *= p
	}

	return product
}

func Permutation(n, k int) int {
	numerator := make([]int, 0)
    product := 1
	
	if k == 0 {
		return 1
	}

	v1 := n
	for i:=0; i<=n; i++ { //Stores a slice of all the values to multiply
        //Because I need the i to incrment in order for the loop to run, i keep "i++", but I will stop it manually here
        if i > 1 {
            v1--
            numerator = append(numerator, v1)
        } else { //only happens when i = 0
            v1 -= i
            numerator = append(numerator, v1)
        }
	}

	for i:=0; i<k; i++ { //wanna get the first v2 numbers
		product *= numerator[i]
	}

	return product
}

// Function to compute n choose k
func Combination(n, k int) int {
	if k == 0 { //dont think if k is 0
		return 1
	} else if n == k { //dont think if they are equal
		return 1
	} else if n-1 == k { //dont think if they are one off
		return n 
	} else {
		if k > n/2 { //if the number is too big for the computer, reduce it here.
			v1 := n-k
			answer := 1
			for i := 1; i <= v1; i++ { //I copied the code here, because it didn't work when i tried running Combination(n, v1)
				answer *= n
				n--
				answer /= i
			}
			return answer
		}

		answer := 1
		for i := 1; i <= k; i++ {
			answer *= n
			n--
			answer /= i
		}
		return answer
	}
}

//It works with small numbers

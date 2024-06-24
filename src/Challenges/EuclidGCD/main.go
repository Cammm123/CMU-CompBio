package main

import (
	"fmt"
	"time"
	"log"
)

func main() {
	fmt.Println("Creating a fucntion to complete Euclid's GCD Theorem")
	fmt.Println(GCD(63, 42))
	fmt.Println(TrivialGCD(63, 42))

	x:=37820268014
	y:=27314794329

	//timing TrivialGCD
	start := time.Now() //start stopwatch
	TrivialGCD(x,y)
	elapsed := time.Since(start) //stopping stopwatch

	log.Printf("Trivial GCD took %s", elapsed) //printing to console in a pretty way

	//timing EuclidGCD
	start2:=time.Now()
	GCD(x, y)
	elapsed2 := time.Since(start2)

	log.Printf("EuclidGCD took %s", elapsed2)

}

func GCD(a, b int) int {

	for a != b {
		if a > b {
			a -= b
		} else if b > a {
			b -= a
		}
	}
	return a //Can be b as well since they are now the same
}

func Min2(a, b int) int {
	min := 0
	if a > b {
		min = b
	} else if b > a {
		min = a
	}

	return min
}

func TrivialGCD(a, b int) int {
	d:=0

	m:= Min2(a,b)
	for p:=1;p<=m;p++ {
		if a%p ==0 && b%p ==0 { //&& stands for and
			//if we are here, it must be a divisor for both
			d = p
		}
	}

	return d
}
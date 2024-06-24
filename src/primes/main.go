package main

import (
	"fmt"
	"math"
	"log"
	"time"
)

func main() {
	fmt.Println("Finding Primes!")

	n:=100000000

	start := time.Now()
	TrivialPrimeFinder(n)
	elapsed := time.Since(start)
	log.Printf("TrivialPrimeFinder took %s", elapsed)

	start2 := time.Now()
	SieveOfEratosthenes(n)
	elapsed2 := time.Since(start2)
	log.Printf("Sieve of Eratosthenes took %s", elapsed2)

	//fmt.Println(TrivialPrimeFinder(n))
	//fmt.Println(SieveOfEratosthenes(n))
	//fmt.Println(ListPrimes(n))
}

/*
TrivialPrimeFinder funtion takes an integer as input
and returns a slice of boolean variable storing the primality
of each non-negative integer up to and including n.
*/

func TrivialPrimeFinder(n int) []bool {
	primeBooleans := make([]bool, n+1)

	for p:=2;p<=n;p++ { //default false values. Excluding 0 and 1
		if IsPrime(p) == true { //can make it more simple if you say: primeBooleans[p] = IsPrime(p)
			primeBooleans[p] = true
		}
	}

	return primeBooleans
}

/*
IsPrime function will take integer as input
and return true if prime, false otherwise
*/
func IsPrime(p int) bool {
	for k:=2; float64(k)<=math.Sqrt(float64(p)); k++ {
		//is k a divisor of p? If so, return false
		if p%k == 0 {
			return false
		}
	}
	//surviving all of these checks means that p is prime
	return true
}

/*
----------------------------Sieve of Eratosthenes---------------------------------
*/

func SieveOfEratosthenes(n int) []bool {
	primeBooleans := make([]bool, n+1)
	//everyone starts as false by default
	//now, srt everything from 2 onward equal to true
	for p:=2; p<=n; p++ {
		primeBooleans[p] = true
	}

	for p:=2; float64(p)<=math.Sqrt(float64(n)); p++ {
		//is p prime? If so, cross of its multiples
		if primeBooleans[p] == true {
			primeBooleans = CrossOffMultiples(primeBooleans, p)
		}
	}

	return primeBooleans
}
/*Functions CrossOffMultiples() will take int p and boolean slice of prime booleans,
	and return an updated slice with all multiples of p are set to false.
	Or, all variables in the array whose indices are mutiples of p
*/

func CrossOffMultiples(primeBooleans []bool, p int) []bool {
	n := len(primeBooleans) - 1

	//conside every mutiple of p, starting with 2p, 
	//and "cross it off" by setting its corresponding entry of the slice to false

	for k:=2*p; k<=n; k+=p {
		primeBooleans[k] = false
	}

	return primeBooleans
}

func ListPrimes(n int) []int {
	primeList := make([]int, 0)

	primeBooleans := SieveOfEratosthenes(n)

	for p:=range primeBooleans{
		if primeBooleans[p] == true {
			primeList = append(primeList, p)
		}
	}

	return primeList
}
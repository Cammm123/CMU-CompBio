package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("HW 0!")
	fmt.Println(Permutation(12, 1))
	fmt.Println(Combination(10, 6))
	fmt.Println(Power(-1, 3))
	a := make([]int, 3)
	a[0] = 6
	a[1] = 9
	a[2] = 16
	fmt.Println(GCDArray(a))
	//fmt.Println(NextPerfectNumber(8128))
	fmt.Println(ListMersennePrimes(61))
}

/* -------------------------------2.1---------------------------------*/
func Combination(n, k int) int {
	if n==0 {
		panic("Error: Zero was inputed for the number of items in Combination().")
	}

	PermutationStatistic := Permutation(n, k)
	CombinationStatistic := PermutationStatistic / Factorial(k)

	return CombinationStatistic
}

func Permutation(n, k int) int {
	if n == 0 {
		panic("Error: Zero was inputed for the number of items in Permutation().")
	}

	numerator := n 
	denominator := n-k

	PermutationStatistic := Factorial(numerator) / Factorial(denominator)

	return PermutationStatistic
}

func Factorial(n int) int {

	product := 1

	for p:=1; p<=n; p++ {
		product *= p
	}

	return product
}

/* -------------------------------2.2---------------------------------*/

func Power(n, k int) int {
	power := 0 //place holder
	if k == 0 {
		power = 1

	} else if k < 0 {
		product := 1
		
		for i:=1; i<=k; i++ {
			product *= n
		}
		power = 1/product

	} else if k > 0 {
		product := 1

		for i:=1; i<=k; i++ {
			product *= n
		}
		power = product
	}
	
	return power
}

func SumProperDivisors(n int) int {
	sum := 0

	for i:=1; i<n; i++ {
		if n%i == 0 {
			sum += i
		}
	}
	return sum
}

/* -------------------------------2.3---------------------------------*/

func FibonacciArray(n int) [] int {
	
	sequence := make([]int, n+1)
	if n == 0 {
		sequence[0] = 1
	} else if n==1 {
		sequence[0] = 1
		sequence[1] = 1
	}

	for i:=2; i<=n; i++ {
		sequence[0] = 1
		sequence[1] = 1
		next := sequence[i-2] + sequence[i-1] //adding the last two numbers
		sequence[i] = next
		}
	
	return sequence
}

func SumIntegerArray(list []int) int {
	value := 0
	for i:=range list {
		value += list[i]
	}

	return value
}

func SumIntegers(numbers ...int) int {
	return SumIntegerArray(numbers)
}

func DividesAll(a []int, d int) bool {
	list := a

	if d == 0 {
		return false
	} else {
		for i:=range list {
			if list[i] % d != 0 {
				return false
			}
		}
	}

	return true
}

func MaxIntegerArray(list []int) int {
	max := list[0]

	for i:=range list {
		if list[i] > max {
			max = list[i]
		}
	}

	return max
}

func MaxIntegers(numbers ...int) int {
	return MaxIntegerArray(numbers)
}

func MinIntegerArray(list []int) int {
	min := list[0]

	for i:=range list {
		if list[i] < min {
			min = list[i]
		}
	}

	return min
}

func MinIntegers(numbers ...int) int {
	return MinIntegerArray(numbers)
}


func GCDArray(a []int) int {
	IntegerNumbers := a
	MaxInt := MaxIntegerArray(IntegerNumbers)
	MinInt := MinIntegerArray(IntegerNumbers)
	divisor := MinInt
	
	
	primeList := ListPrimes(MaxInt) //Using MaxInt because there is no way that we need a higher number of primes

	for i := 0; primeList[i] <= MinInt; i++ {
		if DividesAll(IntegerNumbers, primeList[i]) == false {
			continue
		} else {
			divisor = primeList[i]
			break
		}
	}
	//If the above test failed, I will refresh and try here
	
	if DividesAll(IntegerNumbers, divisor) == false {
		if divisor != 0 {
			divisor = 0
		}

		for divisor <= MaxInt {
			if DividesAll(IntegerNumbers, divisor) == false {
				divisor += MinInt
			}
		}
	}
	//If eveything failed, just set it to 1
	if DividesAll(IntegerNumbers, divisor) != true {
		divisor = 1
	}

	return divisor
}

//Split into prime factorization and multiply the common ones together
//I can do this by making a list of prime numbers
//this list can be set to a particular length depending on what the highest number is
//the last number in ListOfPrimes is under the MaxInt.
//I can calculate that by iterating through a massive list of primes and comparing it to the MaxInt

//My FinalListOfPrimes gets set to all the IntegerNumbers common primes.
//I find that by running a Condensor to condense the primes into the shared ones
//The Condensor iterates through each number's PrimeFactorization and only takes the shared one
//PrimeFactorization works by repetedly dividing by 2 until the number isn't divisible any more
//	at that point, I search for a prime in ListOfPrimes to divide it by. Once the final number
//	reaches 1, I know to stop.
//

/* -------------------------------2.4---------------------------------*/

func NewSumProperDivisors(n int) int {
	sum := 0
	list_of_factors := make([]int, 0)
	var new_factor int

	for i:=1; i<(n/2); i++ {
		if n%i == 0 {
			list_of_factors = append(list_of_factors, i)
			sum += i
		}
	}

	for i:=range list_of_factors {
		new_factor = n/list_of_factors[i]
		if new_factor >= n/2 {
			sum += new_factor //only add on if you are actually new
		}
		
	}
	
	//This all works, but because 1 is a factor, I need to cancel n
	sum -= n

	return sum
}

func IsPerfect(n int) bool {
	if NewSumProperDivisors(n) == n {
		return true
	} else {
		return false
	}
}

func NextPerfectNumber(n int) int {
	var next int

	if n < 8128 {
		if n < 6 {
			next = 6
            return next
		} else if n >= 6 && n < 28{
			next = 28
            return next
		} else if n >= 28 && n < 496 {
			next = 496
            return next
		} else if n >= 496 && n < 8128 {
			next = 8128
            return next
		}
	}

	for i:=8; i<=10; i+=2 { //set limit to 10 because that is just way too big
		if n > 2^i * (2^(i+1) - 1) {
			panic("Error: That is too big of a number")
		}
		
		next = 2^i * (2^(i+1) - 1)
	}
	
	return next
}

func NewPower(n, k int) int {
	power := 0 //place holder
	if k == 0 {
		power = 1

	} else if k < 0 {
		product := 1
		
		for i:=1; i<=k; i++ {
			product *= n
		}
		power = 1/product

	} else if k > 0 {
		product := 1

		for i:=1; i<=k; i++ {
			product *= n
		}
		power = product
	}
	
	return power
}

func ListMersennePrimes(n int) []int {
	array := make([]int, 0)
	for p:=2; p<=n; p++ { //start at 2 to avoid the 1
		next := NewPower(2, p)
		next--
		if IsPrime(next) == true {
			array = append(array, next) //only add if prime
		}
	}

	return array
}

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

/* -------------------------------2.5---------------------------------*/

func NextTwinPrimes(n int) (int, int) {
	primeList := ListPrimes(1000) //I dont think that we will go more than the 100 first primes
    var integer1 int
    var integer2 int
	for i:=range primeList {
		if n < primeList[i] {
			next:= i+1
			if primeList[i] - primeList[next] == -2 {
				integer1 = primeList[i]
            	integer2 = primeList[next]
            	break
			} else {
				continue //keep going through the loop
			}
		}
	}
	
	return integer1, integer2
}

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

/* -------------------------------4.1---------------------------------*/

func Richness(sample map[string]int) int {
	count := 0 //count of animals
	for i := range sample {
		if sample[i] > 0 {
			count++
		}
	}
	return count
}

func SumOfValues(sample map[string]int) int {
	sum := 0
	for i := range sample {
		sum += sample[i]
	}

	return sum
}

func SimpsonsIndex(sample map[string]int) float64 {
	var index float64
	denominator := SumOfValues(sample)

	for i := range sample {
		index += math.Pow(float64(sample[i])/float64(denominator), 2)
	}

	return index
}

/*
func Search(sample map[string]int) int {
	var max int
	for i := range sample {
		if i==0 || sample[i] > max {
			max = sample[i]
		}
	}
	return max
}
*/

/* -------------------------------4.2---------------------------------*/

/*
Sudo Code:

BrayCurtisDistance(sample1, sample2)
total1 ← SumOfValues(sample1)
total2 ← SumOfValues(sample2)
average ← (total1 + total2)/2.0
sum ← SumOfMinima(sample1, sample2)
return 1 - sum/average
*/

func BrayCurtisDistance(sample1, sample2 map[string]int) float64 {
	total1 := SumOfValues(sample1)
	total2 := SumOfValues(sample2)
	average := float64(total1 + total2) / 2.0 //sets to float64
	sum := SumOfMinima(sample1, sample2)
	
	return 1 - float64(sum)/float64(average)
}

func SumOfMinima(sample1, sample2 map[string]int) int {
	var sum int
	var count int //I am making a variable to count the amount of times we only add on val1 or val2.
	//If count gets too high, I can assume that each item is separate and return 0

	// Iterate over the keys of the first map
	for key, val1 := range sample1 {
		val2, exists := sample2[key] //exists is a boolean
		if exists { //if the key exists for sample2 as well,
			sum += Min2(val1, val2) //add on the minimum between them
		}
	}

	// Iterate over the keys of the second map that are not in the first map
	for key, val2 := range sample2 {
		if _, exists := sample1[key]; !exists { //check if key exists in sample1 as well. if it doesnt,

			sum += val2 //add on val2, otherwise, we take care of it in first for loop
			count += 1
			
		}
	}

	//Check if count was too high
	if count >= len(sample1) + len(sample2) {
		return 0
	}
	return sum
}

// Note: for the sake of convenience, we are providing a Min2() function below.
func Min2(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func SumOfMaxima(sample1, sample2 map[string]int) int {
	var sum int

	// Iterate over the keys of the first map
	for key, val1 := range sample1 {
		val2, exists := sample2[key] //exists is a boolean
		if exists { //if the key exists for sample2 as well,
			sum += Max2(val1, val2) //add on the minimum between them
		} else { //just add val1
            sum += val1
        }
	}

	// Iterate over the keys of the second map that are not in the first map
	for key, val2 := range sample2 {
		if _, exists := sample1[key]; !exists { //check if key exists in sample1 as well. if it doesnt,

			sum += val2 //add on val2, otherwise, we take care of it in first for loop
		}
	}
	
    return sum
}

func Max2(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func JaccardDistance(sample1, sample2 map[string]int) float64 {
	sum1 := SumOfMinima(sample1, sample2)
	sum2 := SumOfMaxima(sample1, sample2)

	return 1 - float64(sum1)/float64(sum2)
}
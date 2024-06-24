package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
)

func main() {
	fmt.Println("Homework 1")
	//fmt.Println(RelativelyPrimeProbability(1, 1000000000000, 1000000000))
	fmt.Println(SharedBirthdayProbability(10000, 100000))
	fmt.Println(GenerateMiddleSquareSequence(23, 2))

	count := 0

	for seed := 1000; seed < 9999; seed++ {
		sequence := GenerateMiddleSquareSequence(seed, 4)
		periodLength := len(sequence) - 1 // period length is sequence length - 1
		if periodLength <= 10 {
			count++
		}
	}

	fmt.Println("Number of seeds with period <= 10:", count)

	fmt.Println("Let a = 5, c = 1, and m = 213 = 8192. What is the period of the linear congruential generator when the seed is equal to 1?:", len(GenerateLinearCongruenceSequence(1, 5, 1, 8192)))

}

/* -------------------------------6.1---------------------------------*/

func WeightedDie() int {
	x := rand.Intn(2) //random number between 0 and 1
	var answer int

	if x == 0 {
		answer = 3
	} else if x == 1 {
		answer = Random()
	} else {
		panic("Error: There has been a problem when running the dice")
	}

	return answer
}

// The not three
func Random() int {
	x := rand.Intn(5) //random number bewteen 0 and 4

	if x == 0 {
		return 1
	} else if x == 1 {
		return 2
	} else if x == 2 {
		return 4
	} else if x == 3 {
		return 5
	} else if x == 4 {
		return 6
	} else {
		panic("Error: there was an error with random")
	}
}

/*
Square with side lengths 2
Inside square is a circle radius 1
 --> area of square is 2^2
 --> area of circle is pi(1^2)
probability of selecting a point inside that circle is pi/4
 --> 4 * probability = pi
*/

func EstimatePi(numPoints int) float64 {
	var countInsideCircle int

	for i := 0; i < numPoints; i++ {
		// Generate random point within the square of side length 2
		x := (rand.Float64() * 2) - 1 // random float in [-1, 1)
		y := (rand.Float64() * 2) - 1 // random float in [-1, 1)

		// Check if the point is inside the circle of radius 1
		if x*x+y*y <= 1 {
			countInsideCircle++
		}
	}

	// Estimate pi
	return 4 * float64(countInsideCircle) / float64(numPoints)
}

/* -------------------------------6.2---------------------------------*/

func EuclidGCD(a, b int) int {
	for a != b {
		if a > b {
			a = a - b
		} else {
			b = b - a
		}
	}
	return a
}

func RelativelyPrime(a, b int) bool {
	//quickly check if any of the numbers are 0
	if a == 0 {
		return true
	} else if b == 0 {
		return true
	} else { //now actually start the code
		if EuclidGCD(a, b) == 1 {
			return true
		} else {
			return false //otherwise not relatively prime
		}
	}
}

/*
	 simulate a ton of running
		collect how often its prime or not
		return probability
*/
func RelativelyPrimeProbability(lowerBound, upperBound, numPairs int) float64 {
	countYes := 0
	countTotal := 0

	for i := 0; i < numPairs; i++ { //go through ten thousand times for good probability
		a := (rand.Intn(upperBound-lowerBound+1) + lowerBound)
		b := (rand.Intn(upperBound-lowerBound+1) + lowerBound)

		if RelativelyPrime(a, b) {
			countYes++
		}
		countTotal++

	}

	probability := float64(countYes) / float64(countTotal)
	return probability
}

/* -------------------------------6.3---------------------------------*/

func HasRepeat(a []int) bool {
	count := 0
	for i := range a {
		check := a[i]
		for u := range a {
			if check == a[u] {
				count++ //if they are the same, add to count
			}
		}
		if count > 1 {
			return true
		} else {
			count = 0 //otherwise, set it back to 0
		}
	}

	return false //otherwise, if you can never find a pair, false.
}

func SimulateOneBirthdayTrial(num_people int) bool {
	a := make([]int, 0)

	//generate random number from 1-365 and add to room
	for i := 0; i < num_people; i++ {
		x := rand.Intn(365) + 1
		a = append(a, x)
	}

	//now check if there is a repeat
	return HasRepeat(a)
}

func SharedBirthdayProbability(numPeople, numTrials int) float64 {
	countYes := 0
	countTotal := 0
	for i := 0; i < numTrials; i++ {
		if SimulateOneBirthdayTrial(numPeople) {
			countYes++
		}
		countTotal++
	}

	probability := float64(countYes) / float64(countTotal)
	return probability
}

func TakeZero(x int) int {
	x /= 10
	return x
}

func CountNumDigits(x int) int {
	if x == 0 {
		return 1
	}
	return int(math.Log10(float64(x))) + 1
}

func Pow10(x int) int {
	a := 1
	for i := 0; i < x; i++ {
		a *= 10
	}
	return a
}

func SquareMiddle(x, numDigits int) int {
	if numDigits%2 != 0 || x < 0 || numDigits <= 0 || CountNumDigits(x) > numDigits {
		return -1
	}

	target := 2 * numDigits
	newNum := x * x

	// Convert newNum to a string
	strNum := strconv.Itoa(newNum)

	// Add zeros to the string to ensure it has at least target digits
	for len(strNum) < target {
		strNum = "0" + strNum
	}

	// Extract the middle numDigits from the padded string
	start := (len(strNum) - numDigits) / 2
	answer, _ := strconv.Atoi(strNum[start : start+numDigits])

	return answer
}

func GenerateMiddleSquareSequence(seed, numDigits int) []int {
	sequence := make([]int, 0)
	seen := make(map[int]bool) //according to online, map is faster than checking previous using variable

	// Append the initial seed to the sequence and mark it as seen
	sequence = append(sequence, seed)
	seen[seed] = true

	current := seed
	for i := 1; i < 10; i++ { // Limiting to 10 iterations, I can adjust this as needed
		newNum := SquareMiddle(current, numDigits)
		if newNum == -1 {
			break
		}
		if seen[newNum] { //is it in the map?
			sequence = append(sequence, newNum)
			break // Exit loop if newNum has been seen before
		}
		seen[newNum] = true
		sequence = append(sequence, newNum)
		current = newNum
	}

	return sequence
}

func ComputePeriodLength(a []int) int {
	seen := make(map[int]int)
	var answer, index1, index2 int
	found := false

	for i, num := range a {
		if idx, ok := seen[num]; ok {
			index1 = idx
			index2 = i
			found = true
			break
		} else {
			seen[num] = i
		}
	}

	if found {
		answer = index2 - index1
		if answer < 0 {
			answer *= -1
		}
		return answer
	}

	return 0
}

func GenerateLinearCongruenceSequence(seed, a, c, m int) []int {
	sequence := make([]int, 0)
	seen := make(map[int]bool)
	current := seed

	//add the current seed
	sequence = append(sequence, current)
	seen[current] = true

	for { //repeat until break aka repeat
		next := LinearCongruenceSequenceOnce(current, a, c, m)
		if seen[next] {
			sequence = append(sequence, next) //we want that one repeaet
			break                             //then we leave
		}
		//otherwise
		sequence = append(sequence, next)
		seen[next] = true
		current = next
	}

	return sequence
}

func LinearCongruenceSequenceOnce(current, a, c, m int) int {
	next := (a*current + c) % m
	return next
}

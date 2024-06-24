package main

import (
	"fmt"
	"math/rand"
	//"time"
)

func main() {
	fmt.Println("Rolling Dice!")

	numTrials := 10000000
	fmt.Println("Estimated houseedge with", numTrials, "trials is: ", ComputeCrapsHouseEdge(numTrials))
}

/*
RollDie()
Input: none
Output: a pseudorandom integer between 1 and 6 incluseively
*/

func RollDie() int {
	x := rand.Intn(6) //x is between 0 and 5
	return x + 1
}

/*
SumTwoDice()
input:none
output: the simulated sum of two dice (between 2 and 12)
*/
func SumTwoDie() int {
	roll1 := RollDie()
	roll2 := RollDie()

	return roll1 + roll2
}

/*
SumDice() = simulates the process of summing n dice
Input: an integer numDice
Output: the sum of numDice simulated dice
*/
func SumDice(numDice int) int {
	sum := 0

	for i:=0; i<numDice; i++ {
		roll := RollDie()
		sum += roll
	}

	return sum
}

/*
PlayCrapsOnce() simulates one game of craps.
Input: none
Output: true if the game is a win and false if its a loss
*/
func PlayCrapsOnce() bool {
	firstRoll := SumDice(2)
	if firstRoll == 7 || firstRoll == 11 {
		//winner
		return true
	} else if firstRoll == 2 || firstRoll == 3 || firstRoll == 12 {
		//lose
		return false
	} else {
		// keep rolling until you hit a 7 or original roll
		for {
			newRoll := SumDice(2)
			if newRoll == firstRoll {
				//winner!
				return true
			} else if newRoll == 7 {
				//loser
				return false
			}
		}
	}
}

/*
ComputeCrapsHouseEdge() estimates the "house edge" of craps over multiple simulations
Input: an integer correpsonding to the number of simulations we want to run
Output: house edge of craps (average amount won or lost over the number of simulations per unit of currency)
*/
func ComputeCrapsHouseEdge(numTrials int) float64 {
	count := 0 //will keep track of amount won (+) or lost (-)

	//run n trials and update count accordingly
	for i:=0; i<numTrials; i++ {
		//play game
		outcome := PlayCrapsOnce()
		if outcome{ //also can say "if outcome == true"
			//winner
			count++
		} else {
			//loser
			count--
		}
	}

	return float64(count)/float64(numTrials)
}
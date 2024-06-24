package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Elections!")

	electoralVoteFile := "data/electoralVotes.csv"
	pollFile := "data/debates.csv"

	//now, read them in and store as maps
	electoralVotes := ReadElectoralVotes(electoralVoteFile)
	polls := ReadPollingData(pollFile)

	numTrials := 1000000
	marginOfError := 0.1

	probability1, probability2, probabilityTie := SimulateMultipleElections(polls, electoralVotes, numTrials, marginOfError)
	fmt.Println("Probability of candidate 1 winning: ", probability1)
	fmt.Println("Probability of candidate 2 winning: ", probability2)
	fmt.Println("Probability of tie: ", probabilityTie)
}

/* 
SimulateMultipleElections() runs a Monte Carlo Simulation with mutiple trials to simulate a 
	presidential elections
Input: polling data as a map of states to percentages for candidate 1, 
	a map of state names to Electoral College votes,
	an integer corresponding to the number of trials we want to run,
	and a decimal margin of error
Output: Three probabilities corresponding to the likelihood of each of two candidates winning or a tie.
*/
func SimulateMultipleElections(polls map[string]float64, electoralVotes map[string]uint, numTrials int, marginOfError float64) (float64, float64, float64) {
	winCount1 := 0
	winCount2 := 0
	tieCount := 0

	//simulate a single election n times and update count each time
	for i:=0; i<numTrials; i++ {
		votes1, votes2 := SimulateOneElection(polls, electoralVotes, marginOfError)

		//who won?
		if votes1 > votes2 {
			winCount1++
		} else if votes2 > votes1 {
			winCount2++
		} else {
			tieCount++
		}
	}

	//divide number of wins by number of trials
	probability1 := float64(winCount1)/float64(numTrials)
	probability2 := float64(winCount2)/float64(numTrials)
	probabilityTie := float64(tieCount)/float64(numTrials)

	return probability1, probability2, probabilityTie
}

/*
SimulateOneElection simulates an election from polling data
Input: polling data as a map of states to percentages for candidate 1, 
	a map of state names to Electoral College votes,
	and a decimal margin of error
Output: the number of Electoral College votes in a simulated election 
	for each of candidate 1 and candidate 2
*/
func SimulateOneElection(polls map[string]float64, electoralVotes map[string]uint, marginOfError float64) (uint, uint) {
	var collegeVotes1 uint
	var collegeVotes2 uint

	//range over each state, and simulate the election each one.
	for state, pollingValue := range polls {
		//first lets grab the number of electoral college votes
		numVotes := electoralVotes[state]

		//lets adjust the polling value with some noise
		adjustedPoll := AddNoise(pollingValue, marginOfError)

		//who won the state? (based on adjusted number)
		if adjustedPoll >= 0.5 { //remember this is for candidate1, so if its bigger than 0.5, they won
			collegeVotes1 += numVotes
		} else { //otherwise, candidate2 won
			collegeVotes2 += numVotes
		}

	}

	return collegeVotes1, collegeVotes2
}

/*
AddNoise() adjusts a polling percentage based on some randomization sampled from a normal distribution.
Input: two decimals, a ;olling value and a margin of error.
Output: a decimal coresponding to the adjusted polling value.

*/
func AddNoise(pollingValue, marginOfError float64) float64 {
	x := rand.NormFloat64()
	//x has a 95% chance of being between -2 and 2
	x /= 2.0
	//x has a 95% chance of being between -1 and 1

	x *= marginOfError
	//now x is in place

	//want: x to have a 95% chace of being between -marginOfError and +marginOfError
	return x + pollingValue
}

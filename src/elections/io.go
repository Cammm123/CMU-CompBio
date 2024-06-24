package main

import  (
	"os"
	"strconv"
	"strings"
)

/*
ReadElectoralVotes() processes the number of ectoral votes for each state
Input: a filename string.
Output: a map that associates each state name (string) to an unsigned ineger corresponding
	to its number of Electoral College Votes
*/
func ReadElectoralVotes(filename string) map[string]uint {
	electoralVotes := make(map[string]uint)

	//read in the file contents
	fileContents, err := os.ReadFile(filename) //fileContents is a slice of bytes
	Check(err)

	giantString := string(fileContents)
	//splt string into lines
	lines := strings.Split(giantString, "\n") // \n means new line

	//range over lines, parse each line, and add values to our map
	for _, currentLine := range lines {
		lineElements := strings.Split(currentLine, ",")
		//lineElements has two items: State name and number of electoral votes (as a string)
		stateName := lineElements[0]

		//parse the number of electoral votes
		numVotes, err := strconv.Atoi(lineElements[1])
		Check(err)

		//convert to uint and place into map
		electoralVotes[stateName] = uint(numVotes)
	}

	return electoralVotes
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

/*
ReadPollingData() will parse polling percentqages froma file
Input: a filename string
Output: A map of state names (strings) to floats corresponding to the percentages of candidate 1
func ReadPollingData(filename string) map[string]float64 {
*/
func ReadPollingData(filename string) map[string]float64 {
	candidate1Percentages := make(map[string]float64)

	fileContents, err := os.ReadFile(filename)
	Check(err)
	//convert to string
	giantString := string(fileContents)

	lines := strings.Split(giantString, "\n")

	//range over each line of the file and parse in the data
	for _, currentLine := range lines {
		//splt the current line into its elemnts
		lineElements := strings.Split(currentLine, ",")
		//lineElements has three items (state name and two percentages)
		stateName := lineElements[0]
		percentage1, err := strconv.ParseFloat(lineElements[1], 64)

		Check(err)

		//normalize percentage (divide by 100) and set the appropriate map value
		candidate1Percentages[stateName] = percentage1/100.0
	}

	return candidate1Percentages
}
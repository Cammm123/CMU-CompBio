package main

import "sort"

type PairwiseAlignment [2]string

//CopyGraph returns a copy of the input graph.
func CopyGraph(graph map[string][]string) map[string][]string {
	newGraph := make(map[string][]string)

	for key, value := range graph {
		newArr := make([]string, len(value))
		n := copy(newArr, value)
		if n != len(value) {
			panic("Something wrong happend when copying graph.")
		}
		newGraph[key] = newArr
	}
	return newGraph
}

//MtxEq takes two matrices and returns true if they are equal
//and false otherwise.
func MtxEq(mtx1, mtx2 [][]int) bool {
	if len(mtx1) != len(mtx2) {
		return false
	}
	for i := range mtx1 {
		if len(mtx1[i]) != len(mtx2[i]) {
			return false
		}
		for j := range mtx1[i] {
			if mtx1[i][j] != mtx2[i][j] {
				return false
			}
		}
	}
	return true
}

//GraphEq takes 2 graphs graph1 and graph2.
//It returns if graph1 and graph2 has the same keys and values.
func GraphEq(graph1, graph2 map[string][]string) bool {
	for key1, value1 := range graph1 {
		value2, found := graph2[key1]
		if !found {
			return false
		}
		if !StrSliceEq(value1, value2) {
			return false
		}
	}
	for key2, value2 := range graph2 {
		value1, found := graph1[key2]
		if !found {
			return false
		}
		if !StrSliceEq(value1, value2) {
			return false
		}
	}
	return true
}

//StrSliceEq takes 2 string slices slice1 and slice2.
//It returns if slice1 == slice2.
func StrSliceEq(slice1, slice2 []string) bool {
	sort.Strings(slice1)
	sort.Strings(slice2)
	if len(slice1) != len(slice2) {
		return false
	}
	for i := 0; i < len(slice1); i++ {
		if slice1[i] != slice2[i] {
			return false
		}
	}
	return true
}

//SumLength takes a collection of strings and returns their total length.
func SumLength(patterns []string) int {
	c := 0

	for i := range patterns {
		c += len(patterns[i])
	}

	return c
}


//GlobalAlignment takes two strings, along with match, mismatch, and gap scores.
//It returns a maximum score global alignment of the strings corresponding to these penalties.
func GlobalAlignment(str1, str2 string, match, mismatch, gap float64) PairwiseAlignment {
	backtrack := GlobalBacktrack(str1, str2, match, mismatch, gap)
	// backtrack will compute scores using recurrences and give me all the backtracking pointers as a 2-D array.
	optAlignment := OutputGlobalAlignment(str1, str2, backtrack)
	// this function will use the backtracking info to construct our alignment
	return optAlignment
}


//OutputGlobalAlignment takes two strings and a backtrack matrix of strings. It returns the (presumably optimal) alignment corresponding to this backtracking.
func OutputGlobalAlignment(str1, str2 string, backtrack [][]string) PairwiseAlignment {
	var a PairwiseAlignment

	//start with row and col indices at the sink
	row := len(str1)
	col := len(str2)

	//now, backtrack to the source 

	for row != 0 || col != 0 {
		if backtrack[row][col] == "UP" {
			//take a single symbol from first alignment
			a[0] = string(str1[row-1]) + a[0]
			a[1] = "-" + a[1]
			row--
		} else if backtrack[row][col] == "LEFT" {
			//take a single symbol from second alignment
			a[0] = "-" + a[0]
			a[1] = string(str2[col-1]) + a[1]
			col--
		} else if backtrack[row][col] == "DIAG" {
			//take a single symbol from each alignment
			a[0] = string(str1[row-1]) + a[0]
			a[1] = string(str2[col-1]) + a[1]
			row--
			col--
		} else {
			panic("Error: Backtrack pointers not set correctly.")
		}
	}

	//now we are ready to return the alignment 
	return a 
}

//GlobalBacktrack takes two strings and alignment penalties/rewards.
//It returns a 2-D slice of strings giving backtracking pointers for an optimal
//global alignment of the two strings using these penalties/rewards.
func GlobalBacktrack(str1, str2 string, match, mismatch, gap float64) [][]string {
	if len(str1) == 0 || len(str2) == 0 {
		panic("Blah")
	}
	numRows := len(str1) + 1
	numCols := len(str2) + 1

	backtrack := make([][]string, numRows)
	for i := range backtrack {
		backtrack[i] = make([]string, numCols)
	}

	// grab the scores in the alignment table
	scoreTable := GlobalScoreTable(str1, str2, match, mismatch, gap)
	// 2-D array of float64

	//first, let's set the backtracking pointers of 0-th row and column
	for j := 1; j < numCols; j++ {
		backtrack[0][j] = "LEFT"
	}
	for i := 1; i < numRows; i++ {
		backtrack[i][0] = "UP"
	}

	// traverse rest of scoreTable, and figure out which previous value was used to update each value
	for i := 1; i < numRows; i++ {
		for j := 1; j < numCols; j++ {
			// check what value was used
			if scoreTable[i][j] == scoreTable[i-1][j]-gap {
				// indel
				backtrack[i][j] = "UP"
			} else if scoreTable[i][j] == scoreTable[i][j-1]-gap {
				backtrack[i][j] = "LEFT"
			} else {
				// I should be more precise than this
				backtrack[i][j] = "DIAG"
			}
		}
	}
	return backtrack
}

//GlobalScoreTable takes two strings and alignment penalties. It returns a 2-D array
//holding dynamic programming scores for global alignment with these penalties.
func GlobalScoreTable(str1, str2 string, match, mismatch, gap float64) [][]float64 {
	if len(str1) == 0 || len(str2) == 0 {
		panic("Blah")
	}

	numRows := len(str1) + 1
	numCols := len(str2) + 1

	// initialize scoring matrix -- would be great for subroutine

	scoreTable := make([][]float64, numRows)
	for i := range scoreTable {
		scoreTable[i] = make([]float64, numCols)
	}

	//first, penalize the 0-th row and column as all gaps
	for j := 1; j < numCols; j++ {
		//0-th row
		scoreTable[0][j] = float64(j) * (-gap)
	}
	for i := 1; i < numRows; i++ {
		//0th column
		scoreTable[i][0] = float64(i) * (-gap)
	}

	// now I am ready to range row by row and apply the GA recurrence relation
	for i := 1; i < numRows; i++ {
		for j := 1; j < numCols; j++ {
			//apply the recurrence relation
			upValue := scoreTable[i-1][j] - gap   //indel
			leftValue := scoreTable[i][j-1] - gap //indel
			var diagonalWeight float64
			if str1[i-1] == str2[j-1] { //match!
				diagonalWeight = match
			} else { // mismatch!
				diagonalWeight = -mismatch
			}
			diagValue := scoreTable[i-1][j-1] + diagonalWeight
			scoreTable[i][j] = MaxFloat(upValue, leftValue, diagValue)
		}
	}

	return scoreTable
}

func MaxFloat(nums ...float64) float64 {
	m := 0.0
	// nums gets converted to an array
	for i, val := range nums {
		if val > m || i == 0 {
			// update m
			m = val
		}
	}
	return m
}

//StringSliceEquals compares if 2 string slices contain the same amount of the same patterns.
func StringSliceEquals(patterns1, patterns2 []string) bool {
	//sort first
	sort.Strings(patterns1)
	sort.Strings(patterns2)

	if len(patterns1) != len(patterns2) {
		return false
	}
	for i := range patterns1 {
		if patterns1[i] != patterns2[i] {
			return false
		}
	}
	return true
}
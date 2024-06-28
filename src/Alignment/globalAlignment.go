package main

// Please do not remove package declarations because these are used by the autograder. If you need additional packages, then you may declare them above.

// Alignment is an array of two strings corresponding to the top and bottom
// rows of an alignment of two strings.
type Alignment [2]string

// Insert your GlobalAlignment() function here, along with any subroutines that you need. Please note the subroutines indicated in the problem description that are provided for you.
func GlobalAlignment(str1, str2 string, match, mismatch, gap float64) Alignment {

	scoreTable := GlobalScoreTable(str1, str2, match, mismatch, gap)
	r := len(str1)
	c := len(str2)

	dp := make([][]int, r+1)
	for i := range dp {
		dp[i] = make([]int, c+1)
	}

	var alignment Alignment

	//we add the letter to each one throughout bottom for loop
	row, col := r, c
	for row > 0 && col > 0 {
		up := scoreTable[row-1][col] - gap
		left := scoreTable[row][col-1] - gap
		diag := scoreTable[row-1][col-1]

		if str1[row-1] == str2[col-1] {
			diag = diag + match
		} else if str1[row-1] != str2[col-1] {
			diag = diag - mismatch
		}
		max := MaxFloats(up, left, diag)

		//match case
		if str1[row-1] == str2[col-1] {
			diag = diag + match
			alignment[0] = string(str1[row-1]) + alignment[0]
			alignment[1] = string(str2[col-1]) + alignment[1]
			//move after adding values
			row--
			col--
		} else if up == max { //update string2, dash for string1

			alignment[0] = string(str1[row-1]) + alignment[0]
			alignment[1] = "-" + alignment[1]

			//move after adding values
			row--
		} else if left == max { //update string1, dash for string2

			alignment[0] = "-" + alignment[0]
			alignment[1] = string(str2[col-1]) + alignment[1]

			//move after adding values
			col--
		} else {
			diag = diag - mismatch
			alignment[0] = string(str1[row-1]) + alignment[0]
			alignment[1] = string(str2[col-1]) + alignment[1]

			row--
			col--
		}
	}

	for row > 0 {

		alignment[0] = string(str1[row-1]) + alignment[0]
		alignment[1] = "-" + alignment[1]
		row--
	}

	for col > 0 {
		alignment[0] = "-" + alignment[0]
		alignment[1] = string(str2[col-1]) + alignment[1]

		col--

	}

	//alignment[0] = newString1
	//alignment[1] = newString2

	return alignment
}

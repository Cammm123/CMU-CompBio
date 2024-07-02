package main

//LocalAlignment takes two strings, along with match, mismatch, and gap scores.
//It returns a maximum score local alignment of the strings corresponding to these penalties.

func LocalAlignment(str1, str2 string, match, mismatch, gap float64) (Alignment, int, int, int, int) {


	numDiags := 0

	scoreTable := LocalScoreTable(str1, str2, match, mismatch, gap)
	r := len(str1)
	c := len(str2)

	dp := make([][]int, r+1)
	for i := range dp {
		dp[i] = make([]int, c+1)
	}

	var alignment Alignment

	//we add the letter to each ne throughout bottom for loop

	var max2 int
    var end1 int
    var end2 int
    var start1 int
    var start2 int
    
    for i := range scoreTable {
        for j := range scoreTable[i] {
            if int(scoreTable[i][j]) > max2 {
                
                max2 = int(scoreTable[i][j])
				end1 = i
				end2 = j
            }
        }
    }
	
	row, col := end1, end2
	for row > 0 && col > 0 && scoreTable[row][col] != 0{

		up := scoreTable[row-1][col] - gap
		left := scoreTable[row][col-1] - gap
		diag := scoreTable[row-1][col-1]

	
        if int(scoreTable[row][col]) == max2 { //this means we reached the max value
            end1 = col
            end2 = row
		}

		if str1[row-1] == str2[col-1] {
			diag = diag + match
			numDiags++
		} else if str1[row-1] != str2[col-1] {
			diag = diag - mismatch
			numDiags++
		}
		max := MaxFloats(up, left, diag, 0)

		//match case
		if up == max { //update string2, dash for string1

			alignment[0] = string(str1[row-1]) + alignment[0]
			alignment[1] = "-" + alignment[1]

			//move after adding values
			row--
		} else if left == max { //update string1, dash for string2

			alignment[0] = "-" + alignment[0]
			alignment[1] = string(str2[col-1]) + alignment[1]

			//move after adding values
			col--
		} else if str1[row-1] == str2[col-1] {
			diag = diag + match
			alignment[0] = string(str1[row-1]) + alignment[0]
			alignment[1] = string(str2[col-1]) + alignment[1]
			//move after adding values
			row--
			col--
		} else {
			diag = diag - mismatch
			alignment[0] = string(str1[row-1]) + alignment[0]
			alignment[1] = string(str2[col-1]) + alignment[1]

			row--
			col--
		}
	}
    
    start1 = col
    start2 = row

	//alignment[0] = newString1
	//alignment[1] = newString2


    return alignment, start2, end2, start1, end1
}
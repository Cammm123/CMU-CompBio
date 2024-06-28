package main

// GlobalScoreTable takes two strings and alignment penalties. It returns a 2-D array
// holding dynamic programming scores for global alignment with these penalties.
func GlobalScoreTable(str1, str2 string, match, mismatch, gap float64) [][]float64 {
	array_2D := make([][]float64, len(str1)+1)

	for row := range array_2D {
		array_2D[row] = make([]float64, len(str2)+1)
	}

	for i := 1; i <= len(str2); i++ {

		array_2D[0][i] = array_2D[0][i-1] - gap

	}

	//subtracts gap every time we index a cell in the row
	for j := 1; j <= len(str1); j++ {
		array_2D[j][0] = array_2D[j-1][0] - gap

	}

	// LOOP THROUHGH ROWS
	for r := 1; r <= len(str1); r++ {
		for c := 1; c <= len(str2); c++ {
			x := array_2D[r][c-1] - gap
			y := array_2D[r-1][c] - gap
			z := array_2D[r-1][c-1]
			//if the substring of first string == substring of the second  string add 1
			if str1[r-1] == str2[c-1] {
				z += match
				array_2D[r][c] = MaxFloats(x, y, z)
				//mismatch case
			} else {
				z -= mismatch
				array_2D[r][c] = MaxFloats(x, y, z)
			}
		}
	}

	return array_2D

}

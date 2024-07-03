package main

//ALL PENALTIES POSITIVE

//ScoreOverlapAlignment takes two strings along with match, mismatch, and gap penalties.
//It returns the maximum score of an overlap alignment in which str1 is overlapped with str2.
//Assume we are overlapping a suffix of str1 with a prefix of str2.
func ScoreOverlapAlignment(str1, str2 string, match, mismatch, gap float64) float64 {
	array_2D := make([][]float64, len(str1)+1)
	for row := range array_2D {
		array_2D[row] = make([]float64, len(str2)+1)
	}
	for i := 1; i <= len(str2); i++ {
		array_2D[0][i] = array_2D[0][i-1] - gap
	}
	// LOOP THROUHGH ROWS
	for r := 1; r <= len(str1); r++ {
		for c := 1; c <= len(str2); c++ {
			x := array_2D[r][c-1] - gap
			y := array_2D[r-1][c] - gap
			z := array_2D[r-1][c-1]
			//if the substring of first string == substring of the second  string add 1
			if str1[r-1] == str2[c-1] {
				z = z + match
				array_2D[r][c] = MaxIntegers(x, y, z)
				//mismatch case
			} else {
				z = z - mismatch
				array_2D[r][c] = MaxIntegers(x, y, z)
			}
		}
	}
	var new_Array []float64
	for x := 0; x < len(str2)+1; x++ {
		new_Array = append(new_Array, array_2D[len(str1)][x])
	}
    //fmt.Println(new_Array)
    //fmt.Println(array_2D)
	return MaxIntegerArray(new_Array)
}
func MaxIntegerArray(list []float64) float64 {
	max := list[0]
	for i := range list {
		if list[i] > max {
			max = list[i]
		}
	}
	return max
}
func MaxIntegers(numbers ...float64) float64 {
	return MaxIntegerArray(numbers)
}

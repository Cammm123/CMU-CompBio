package main

// LCSLength takes two strings as input. It returns the length of a longest common
// subsequence of the two strings.
func LCSLength(str1, str2 string) int {
	if len(str1) == 0 || len(str2) == 0 {
		panic("Empty strings given.")
	}

	array_2D := make([][]int, len(str1)+1)

	for row := range array_2D {
		array_2D[row] = make([]int, len(str2)+1)
	}
	// LOOP THROUHGH ROWS
	for r := 1; r <= len(str1); r++ {
		for c := 1; c <= len(str2); c++ {
			x := array_2D[r][c-1]
			y := array_2D[r-1][c]
			z := array_2D[r-1][c-1]
			//if the substring of first string == substring of the second  string add 1
			if str1[r-1] == str2[c-1] {
				array_2D[r][c] = z + 1
			} else {
				array_2D[r][c] = MaxIntegers(x, y, z)
			}
		}
	}

	return array_2D[len(str1)][len(str2)]
}

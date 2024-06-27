package main

// EditDistance takes two strings as input. It returns the Levenshtein distance
// between the two strings; that is, the minimum number of substitutions, insertions, and deletions
// needed to transform one string into the other.
func EditDistance(str1, str2 string) int {
	array_2D := make([][]int, len(str1)+1)

	for row := range array_2D {
		array_2D[row] = make([]int, len(str2)+1)
	}

	//adding 1 to each cell forward in the first row
	for i := 1; i <= len(str2); i++ {

		array_2D[0][i] = array_2D[0][i-1] + 1

	}

	//adding 1 to each cell down in the first column
	for j := 1; j <= len(str1); j++ {
		array_2D[j][0] = array_2D[j-1][0] + 1

	}

	// LOOP THROUHGH ROWS
	for r := 1; r <= len(str1); r++ {
		for c := 1; c <= len(str2); c++ {
			x := array_2D[r][c-1] + 1
			y := array_2D[r-1][c] + 1
			z := array_2D[r-1][c-1]
			//if the substring of first string == substring of the second  string add 1
			if str1[r-1] != str2[c-1] {
				z++
				array_2D[r][c] = MinIntegers(x, y, z)
			} else {
				array_2D[r][c] = z
			}
		}
	}

	return array_2D[len(str1)][len(str2)]
	//return array_2D
}

func MinIntegerArray(list []int) int {
	min := list[0]

	for i := range list {
		if list[i] < min {
			min = list[i]
		}
	}

	return min
}

func MinIntegers(numbers ...int) int {
	return MinIntegerArray(numbers)
}

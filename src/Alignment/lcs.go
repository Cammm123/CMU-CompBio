// LongestCommonSubsequence takes two strings as input.
// It returns a longest common subsequence of the two strings.
package main

//LongestCommonSubsequence takes two strings as input.
//It returns a longest common subsequence of the two strings

// Please do not remove package declarations because these are used by the autograder. If you need additional packages, then you may declare them above.

// Insert your LongestCommonSubsequence() function here, along with any subroutines that you need. Please note the subroutines indicated in the problem description that are provided for you.
func LongestCommonSubsequence(str1, str2 string) string {

	m := len(str1)
	n := len(str2)

	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if str1[i-1] == str2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	var lcs []string

	//the one that stores the diagonal values.
	//if going to the left and up has the same value, we move diagonally and store that letter
	//if going to the left is better we go to the left (do not store)
	//if going up is better we go up (do not store value)

	i, j := m, n
	for i > 0 && j > 0 {
		if str1[i-1] == str2[j-1] {
			lcs = append(lcs, string(str1[i-1]))
			i--
			j--
		} else if dp[i-1][j] > dp[i][j-1] {
			i--
		} else {
			j--
		}
	}

	//returning the list back to forwards
	//we do this because our LCS list currently holds all the values we need, but in reverse order
	for k := 0; k < len(lcs)/2; k++ {

		lcs[k], lcs[len(lcs)-1-k] = lcs[len(lcs)-1-k], lcs[k]

	}

	return join(lcs)

}

func join(parts []string) string {
	result := ""

	for _, part := range parts {
		result += part
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

package main

import "fmt"

func longestCommonSequence(x, y string) string {
	m := len(x)
	n := len(y)

	lcs := make([][]int, m+1)
	for i := range lcs {
		lcs[i] = make([]int, n+1)
	}

	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ {
			if i == 0 || j == 0 {
				lcs[i][j] = 0
			} else if x[i-1] == y[j-1] {
				lcs[i][j] = 1 + lcs[i-1][j-1]
			} else {
				lcs[i][j] = max(lcs[i-1][j], lcs[i][j-1])
			}
		}
	}

	i, j := m, n
	var lcsString string
	for i > 0 && j > 0 {
		if x[i-1] == y[j-1] {
			lcsString = string(x[i-1]) + lcsString
			i--
			j--
		} else if lcs[i-1][j] > lcs[i][j-1] {
			i--
		} else {
			j--
		}
	}

	return lcsString
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func main() {
	x := "abcd"
	y := "bcad"

	lcs := longestCommonSequence(x, y)
	fmt.Println(lcs)
}

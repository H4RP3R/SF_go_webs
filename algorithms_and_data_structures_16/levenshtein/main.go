// Расстояние Левенштейна, или редакционное расстояние, — метрика cходства между двумя строковыми
// последовательностями. Чем больше расстояние, тем более различны строки. Для двух одинаковых
// последовательностей расстояние равно нулю.

package main

import "fmt"

func levenshteinDistance(s1, s2 string) int {
	m, n := len([]rune(s1)), len([]rune(s2))
	dp := make([][]int, m+1)

	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 0; i <= n; i++ {
		dp[i][0] = i
	}

	for j := 0; j <= n; j++ {
		dp[0][j] = j
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			cost := 0
			if s1[i-1] != s2[j-1] {
				cost = 1
			}
			dp[i][j] = min(dp[i-1][j]+1, dp[i][j-1]+1, dp[i-1][j-1]+cost)
		}
	}

	return dp[m][n]
}

func min(a, b, c int) int {
	if a <= b && a <= c {
		return a
	} else if b <= a && b <= c {
		return b
	} else {
		return c
	}
}

func main() {
	str1 := "Мама мыла раму"
	str2 := "Лена мыла раму"

	distance := levenshteinDistance(str1, str2)
	fmt.Println(distance)
}

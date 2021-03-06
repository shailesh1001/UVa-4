// UVa 10099 - The Tourist Guide

package main

import (
	"fmt"
	"os"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int { return a + b - max(a, b) }

func floydWarshall(n int, matrix [][]int) [][]int {
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	copy(dp, matrix)
	for k := 1; k <= n; k++ {
		for i := 1; i <= n; i++ {
			for j := 1; j <= n; j++ {
				dp[i][j] = max(dp[i][j], min(dp[i][k], dp[k][j]))
			}
		}
	}
	return dp
}

func main() {
	in, _ := os.Open("10099.in")
	defer in.Close()
	out, _ := os.Create("10099.out")
	defer out.Close()

	var n, r, c1, c2, p, s, d, t, kase int
	for {
		if fmt.Fscanf(in, "%d%d", &n, &r); n == 0 {
			break
		}
		matrix := make([][]int, n+1)
		for i := range matrix {
			matrix[i] = make([]int, n+1)
		}
		for r > 0 {
			fmt.Fscanf(in, "%d%d%d", &c1, &c2, &p)
			matrix[c1][c2], matrix[c2][c1] = p, p
			r--
		}
		fmt.Fscanf(in, "%d%d%d", &s, &d, &t)
		dp := floydWarshall(n, matrix)
		trips := t / (dp[s][d] - 1)
		if t%trips != 0 {
			trips++
		}
		kase++
		fmt.Fprintf(out, "Scenario #%d\n", kase)
		fmt.Fprintf(out, "Minimum Number of Trips = %d\n", trips)
	}
}

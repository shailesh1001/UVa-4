// UVa 117 - The Postal Worker Rings Once

package main

import (
	"fmt"
	"math"
	"os"
)

type edge struct {
	n1, n2 string
	l      int
}

func toMatrix(edges []edge) ([]int, int, [][]int) {
	ns := make(map[string]int)
	nl := make(map[string]int)
	idx, sum := 0, 0
	for _, v := range edges {
		sum += v.l
		ns[v.n1]++
		ns[v.n2]++
		if _, ok := nl[v.n1]; !ok {
			nl[v.n1] = idx
			idx++
		}
		if _, ok := nl[v.n2]; !ok {
			nl[v.n2] = idx
			idx++
		}
	}

	matrix := make([][]int, len(nl))
	for i := range matrix {
		matrix[i] = make([]int, len(nl))
		for j := range matrix[i] {
			matrix[i][j] = math.MaxInt32
		}
	}
	for _, v := range edges {
		matrix[nl[v.n1]][nl[v.n2]], matrix[nl[v.n2]][nl[v.n1]] = v.l, v.l
	}

	var odds []int
	for k, v := range ns {
		if v%2 != 0 {
			odds = append(odds, nl[k])
		}
	}
	return odds, sum, matrix
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func floydWarshall(odds []int, matrix [][]int) int {
	for k := 0; k < len(matrix); k++ {
		for i := 0; i < len(matrix); i++ {
			for j := 0; j < len(matrix); j++ {
				matrix[i][j] = min(matrix[i][j], matrix[i][k]+matrix[k][j])
			}
		}
	}
	return matrix[odds[0]][odds[1]]
}

func main() {
	in, _ := os.Open("117.in")
	defer in.Close()
	out, _ := os.Create("117.out")
	defer out.Close()

	var word string
	var edges []edge
	for {
		if _, err := fmt.Fscanf(in, "%s", &word); err != nil {
			break
		}
		if word != "deadend" {
			edges = append(edges, edge{word[:1], word[len(word)-1:], len(word)})
		} else {
			odds, sum, matrix := toMatrix(edges)
			if len(odds) == 0 {
				fmt.Fprintln(out, sum)
			} else {
				fmt.Fprintln(out, sum+floydWarshall(odds, matrix))
			}
			edges = make([]edge, 0)
		}
	}
}

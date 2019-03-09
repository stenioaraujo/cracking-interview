package main

import (
	"fmt"
)

func main() {
	matrix := [][]int{
		{1, 2, 3, 0, 0, 0},
		{4, 5, 6, 0, 0, 0},
		{7, 8, 9, 0, 0, 0},
		{10, 11, 12, 0, 0, 0},
		{10, 11, 12, 0, 0, 0},
		{10, 11, 12, 0, 0, 0},
	}
	printMatrix(matrix)

	matrix = [][]int{
		{1, 2, 3, 1},
		{4, 5, 6, 1},
		{0, 8, 0, 1},
		{10, 11, 12, 1},
	}
	printMatrix(matrix)

	matrix = [][]int{
		{1, 2, 3},
		{4, 0, 6},
		{7, 8, 9},
	}
	printMatrix(matrix)

	matrix = [][]int{
		{1, 2},
		{3, 4},
	}
	printMatrix(matrix)

	matrix = [][]int{
		{1},
	}
	printMatrix(matrix)
}

func zero(matrix *[][]int) {
	m := *matrix

	lines := make([]bool, len(m))
	columns := make([]bool, len(m[0]))

	for i, l := range m {
		for j, e := range l {
			if e == 0 {
				lines[i] = true
				columns[j] = true
			}
		}
	}

	zeroLines(matrix, lines...)
	zeroColumns(matrix, columns...)
}

func zeroLines(matrix *[][]int, lines ...bool) {
	m := *matrix
	for l, shouldZero := range lines {
		if shouldZero {
			for i := 0; i < len(m[0]); i++ {
				m[l][i] = 0
			}
		}
	}
}

func zeroColumns(matrix *[][]int, columns ...bool) {
	m := *matrix
	for c, shouldZero := range columns {
		if shouldZero {
			for i := 0; i < len(m); i++ {
				m[i][c] = 0
			}
		}
	}
}

func printMatrix(matrix [][]int) {
	for _, l := range matrix {
		for _, e := range l {
			fmt.Printf("%d\t", e)
		}
		fmt.Println()
	}
	zero(&matrix)
	fmt.Println("---------------------------------")
	for _, l := range matrix {
		for _, e := range l {
			fmt.Printf("%d\t", e)
		}
		fmt.Println()
	}
	fmt.Println("===================================")
}

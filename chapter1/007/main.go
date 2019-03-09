package main

import (
	"fmt"
	"math"
	"sync"
)

func main() {
	matrix := [][]int32{
		{1, 2, 3, 0, 0, 0},
		{4, 5, 6, 0, 0, 0},
		{7, 8, 9, 0, 0, 0},
		{10, 11, 12, 0, 0, 0},
		{10, 11, 12, 0, 0, 0},
		{10, 11, 12, 0, 0, 0},
	}
	printMatrix(matrix)

	matrix = [][]int32{
		{1, 2, 3, 0},
		{4, 5, 6, 0},
		{7, 8, 9, 0},
		{10, 11, 12, 0},
	}
	printMatrix(matrix)

	matrix = [][]int32{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	printMatrix(matrix)

	matrix = [][]int32{
		{1, 2},
		{3, 4},
	}
	printMatrix(matrix)

	matrix = [][]int32{
		{1},
	}
	printMatrix(matrix)
}

func rotate(matrix *[][]int32) {
	middle := int(math.Round(float64(len(*matrix)/2.0)) - 1)
	var wg sync.WaitGroup
	wg.Add(middle + 1)
	for i := 0; i <= middle; i++ {
		go func(i int) {
			rotateStartEnd(matrix, i, len(*matrix)-1-i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func rotateStartEnd(matrix *[][]int32, start, end int) {
	m := *matrix
	for i := 0; i < end-start; i++ {
		m[start][start+i], m[end-i][start] = m[end-i][start], m[start][start+i]
	}
	for i := 0; i < end-start; i++ {
		m[start+i][end], m[start][start+i] = m[start][start+i], m[start+i][end]
	}
	for i := 0; i < end-start; i++ {
		m[end][end-i], m[start+i][end] = m[start+i][end], m[end][end-i]
	}
}

func printMatrix(matrix [][]int32) {
	for _, l := range matrix {
		for _, e := range l {
			fmt.Printf("%d\t", e)
		}
		fmt.Println()
	}
	rotate(&matrix)
	fmt.Println("---------------------------------")
	for _, l := range matrix {
		for _, e := range l {
			fmt.Printf("%d\t", e)
		}
		fmt.Println()
	}
	fmt.Println("===================================")
}

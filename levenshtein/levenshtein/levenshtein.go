package levenshtein

import (
	"fmt"
)

func LevenshteinDistanceRecurrent(s1, s2 string) int32 {
	if len(s1) == 0 {
		return int32((len(s2)))
	}
	if len(s2) == 0 {
		return int32(len(s1))
	}
	edition := int32(1)
	if s1[0] == s2[0] {
		edition = 0
	}
	//min(substitution, deletion, insertiion)
	return min(edition+LevenshteinDistanceRecurrent(s1[1:], s2[1:]), min(1+LevenshteinDistanceRecurrent(s1[1:], s2), 1+LevenshteinDistanceRecurrent(s1, s2[1:])))
}

func LevenshteinDistance(s1, s2 string, verbose bool) int32 {

	distance := make([][]int32, len(s1)+1, len(s1)+1)

	for i, _ := range distance {
		distance[i] = make([]int32, len(s2)+1, len(s2)+1)
	}
	rows := len(distance)
	columns := len(distance[0])

	for i := 1; i < rows; i++ {
		distance[i][0] = int32(i)
	}

	for j := 1; j < columns; j++ {
		distance[0][j] = int32(j)
	}
	if verbose {
		fmt.Println("Dynamic Programming Matrix with initial values (character position):")
		printMatrix(distance)
		fmt.Println()
	}

	// The loops go from i=0 and j= 0 to i<rows-1 and j<columns respectively instead of i<rows because we add 1 to len(s1) and len(s2) when building the matrix.
	for i := 0; i < rows-1; i++ {
		for j := 0; j < columns-1; j++ {
			edition := int32(1)
			if s1[i] == s2[j] {
				edition = 0
			}
			distance[i+1][j+1] = min(distance[i][j]+edition, min(distance[i][j+1]+1, distance[i+1][j]+1))
		}
	}

	if verbose {
		fmt.Println("Dynamic Programming Matrix after applying levenshtein:")
		printMatrix(distance)
		fmt.Println()
	}
	return distance[rows-1][columns-1]
}

func printMatrix(matrix [][]int32) {
	for i, _ := range matrix {
		fmt.Println(matrix[i])
	}
}

func min(a, b int32) int32 {
	if a < b {
		return a
	}
	return b
}

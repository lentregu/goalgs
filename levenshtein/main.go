package main

import (
	"fmt"

	"github.com/lentregu/goalgs/levenshtein/levenshtein"
)

func main() {
	fmt.Println(levenshtein.LevenshteinDistance("mnopqr", "m1oqp", true))
	fmt.Println(levenshtein.LevenshteinDistance("hola que tal", "ola que tu haces", true))
	fmt.Println(levenshtein.LevenshteinDistanceRecurrent("mnopqr", "m1oqp"))
	fmt.Println(levenshtein.LevenshteinDistanceRecurrent("hola que tal", "ola que tu haces"))
}

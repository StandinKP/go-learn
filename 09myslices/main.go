package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices")

	var fruits = []string{"apple", "banana", "cherry"}
	// fmt.Printf("%T\n", fruits)

	fruits = append(fruits, "orange", "grape")

	// fmt.Println(fruits)

	highScores := make([]int, 5)
	highScores[0] = 100
	highScores[1] = 90
	highScores[2] = 80
	highScores[3] = 70
	highScores[4] = 60

	// fmt.Println(highScores)

	highScores = append(highScores, 50, 23, 442, 23, 4)
	// fmt.Println(highScores)

	sort.Ints(highScores)
	// fmt.Println(highScores)
	// fmt.Println(sort.IntsAreSorted(highScores))

}

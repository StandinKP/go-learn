package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions")
	greet()
	result := add(2, 3)
	fmt.Println("Result is: ", result)
	result2 := proAdd(2, 3, 4, 5, 45, 45, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println("Result is: ", result2)
}

func add(x int, y int) int {
	return x + y
}

func proAdd(values ...int) int {
	sum := 0
	for _, value := range values {
		sum += value
	}
	return sum
}

func greet() {
	fmt.Println("Hello")
}

package main

import "fmt"

func main() {
	fmt.Println("Welcome to arrays")

	var fruits [4]string
	fruits[0] = "Apple"
	fruits[1] = "Orange"
	fruits[3] = "Banana"

	fmt.Println("Fruits:", fruits)

	var vegetables = [2]string{"Carrot", "Potato"}
	fmt.Println("Vegetables:", vegetables)
}

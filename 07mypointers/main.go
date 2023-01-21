package main

import "fmt"

func main() {
	fmt.Println("Welcome to pointers")

	// var ptr *int
	// fmt.Println(ptr)

	myNumber := 23
	var ptr = &myNumber
	fmt.Println(ptr)
	fmt.Println(*ptr)

}

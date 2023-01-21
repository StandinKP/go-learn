package main

import "fmt"

func main() {
	defer fmt.Println("defer")
	defer fmt.Println("defer1")
	defer fmt.Println("defer2")
	fmt.Println("Welcome to ")
}

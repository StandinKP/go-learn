package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	// No inheritance, super, parent in go

	kaushal := User{"Kaushal", 20, "k@k.com", true}
	fmt.Printf("%+v\n", kaushal)
}

type User struct {
	Name   string
	Age    int
	Email  string
	Status bool
}

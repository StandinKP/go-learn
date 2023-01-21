package main

import "fmt"

func main() {
	fmt.Println("If Else in golang")
	loginCount := 10
	var result string

	if loginCount < 10 {
		result = "You are not allowed to login"
	} else if loginCount > 10 {
		result = "You are allowed to login"
	} else {
		result = "10 logins"
	}

	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}

	if num := 9; num < 10 {
		fmt.Println("Smaller")
	} else if num > 10 {
		fmt.Println("Bigger")
	} else {
		fmt.Println("Equal")
	}

}

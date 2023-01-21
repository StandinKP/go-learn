package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to Go!"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter rating for pizza:")

	// comma ok || error ok
	input, _ := reader.ReadString('\n')
	fmt.Println("You entered:", input)
	fmt.Printf("Type of rating is %T \n", input)

}

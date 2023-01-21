package main

import "fmt"

const LoginToken string = "12345"

func main() {
	var username string = "Kaushal"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal int = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float64 = 255.453245345
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	//default values and aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	//implicit type
	var website = "learncodeonline.in"
	fmt.Println(website)
	fmt.Printf("Variable is of type: %T \n", website)

	//no var style
	noOfUsers := 10
	fmt.Println(noOfUsers)
	fmt.Printf("Variable is of type: %T \n", noOfUsers)

	fmt.Println(LoginToken)
}

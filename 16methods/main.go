package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	// No inheritance, super, parent in go

	kaushal := User{"Kaushal", 20, "k@k.com", true}
	fmt.Printf("%+v\n", kaushal)
	kaushal.GetStatus()
	kaushal.newEmail("test@k.com")
	fmt.Printf("%+v\n", kaushal)

}

type User struct {
	Name   string
	Age    int
	Email  string
	Status bool
}

func (u User) GetStatus() {
	fmt.Println("Is the user active:", u.Status)
}

func (u User) newEmail(email string) {
	u.Email = email
	fmt.Println("New email is:", u.Email)
}

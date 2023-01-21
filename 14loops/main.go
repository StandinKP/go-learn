package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in golang")

	days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	fmt.Println("Days of the week:", days)

	// for i, day := range days {
	// 	fmt.Println(i, day)
	// }

	// for i := 0; i < len(days); i++ {
	// 	fmt.Println(i, days[i])
	// }

	roughValue := 1
	for roughValue < 10 {

		if roughValue == 2 {
			goto kp
		}

		if roughValue == 5 {
			roughValue++
			continue
		}

		fmt.Println(roughValue)
		roughValue++
	}

kp:
	fmt.Println("Jump at kp")

}

package main

import "fmt"

func main() {
	fmt.Println("Maps in golang")

	languages := make(map[string]string, 5)

	languages["en"] = "English"
	languages["es"] = "Spanish"
	languages["fr"] = "French"
	languages["de"] = "German"
	languages["it"] = "Italian"

	fmt.Println(languages)
	fmt.Println(languages["en"])

	delete(languages, "en")
	fmt.Println(languages)

	for key, value := range languages {
		fmt.Println(key, value)
	}
}

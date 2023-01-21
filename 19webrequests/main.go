package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("Welcome to web requests")

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%T\n", response)

	defer response.Body.Close() //SHOULD ALWAYS CLOSE THIS

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	content := string(data)
	fmt.Println(content)

}

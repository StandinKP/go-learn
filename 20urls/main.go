package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=gfdgdfgd"

func main() {
	fmt.Println("URLs in go")

	fmt.Println("\nURL: " + myUrl)

	// parsing urls
	result, _ := url.Parse(myUrl)
	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Path)
	// fmt.Println(result.RawQuery)
	// fmt.Println(result.Port())

	// query params
	queryParams := result.Query()
	fmt.Println(queryParams)

}

package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time")

	presentTime := time.Now()
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Mon"))

	createdDate := time.Date(2019, time.March, 1, 0, 0, 0, 0, time.UTC)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 15:04"))
}

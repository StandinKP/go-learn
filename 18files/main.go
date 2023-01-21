package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files in golang")
	content := "Hello World. Put this is file"
	file, err := os.Create("./file.txt")

	checkNilError(err)
	length, err := io.WriteString(file, content)
	checkNilError(err)
	fmt.Printf("Wrote %d bytes\n", length)
	defer file.Close()
	readFile("./file.txt")

}

func readFile(filename string) {
	dataBytes, err := ioutil.ReadFile(filename)
	checkNilError(err)
	fmt.Println("Text: \n", string(dataBytes))
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}

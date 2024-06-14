package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to files in golang!")

	content := "This needs to go inside a file - vegadelalyra.com"

	file, err := os.Create("./mylyragofile.txt")
	checkNilErr(err)

	length, err := io.WriteString(file, content)
	checkNilErr(err)

	fmt.Println("length is:", length)
	defer file.Close()
	fmt.Println(file.Name())
	readFile("./mylyragofile.txt")
}

func readFile(filename string) {
	databyte, err := os.ReadFile(filename)
	checkNilErr(err)

	fmt.Println("Text data inside the file is... \n", string(databyte))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}

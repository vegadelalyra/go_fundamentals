package main

import "fmt"

func main() {
	fmt.Println("Welcome to a class on pointers!")

	var one int // regular var
	fmt.Printf("one: %v\n", one)

	var ptr *int // my first go *int pointer!!!
	fmt.Printf("ptr: %v\n", ptr)

	myNumber := 23
	var ref = &myNumber //& referencing to an existing variable

	fmt.Println("Value of actual pointer is", ref)
	fmt.Println("Pointer of actual pointer is", *ref)

	*ref = *ref * 2
	fmt.Printf("myNumber: %v\n", myNumber)
}

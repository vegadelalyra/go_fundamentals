package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in golang")
	greeter()
	greeterTwo()

	result := adder(3, 5)
	fmt.Println("Result is", result)

	proRes, msg := proAdder(1, 25, 5, 246, 57, 446, 4, 42, 32)

	fmt.Printf("proRes: %v\n", proRes)
	fmt.Printf("msg: %v\n", msg)
}

func adder(i1, i2 int) int {
	return i1 + i2
}

func proAdder(values ...int) (int, string) {
	total := 0

	for _, v := range values {
		total += v
	}

	return total, "Hehehe!"
}

func greeter() {
	fmt.Println("Namstey from golang")
}

func greeterTwo() {
	fmt.Println("Another method")
}

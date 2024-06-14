package main

import (
	"bufio"
	"fmt"
	"os"
)

func main()  {
	welcome := "Welcome to user input!"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our Pizza!")

	// comma ok // err err // comma err syntax
	input, _ := reader.ReadString('\n')
	fmt.Print("Thanks for rating, thing! ", input)
	fmt.Printf("Type of this rating is %T", input)
}
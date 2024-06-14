package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golangs!")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[3] = "Peach"

	fmt.Println("The fruit list is:\n", fruitList)
	fmt.Println("The fruit list is:", len(fruitList))

	var vegList = [3]string{"potato", "beans", "mushroom"}
	fmt.Println("Vegy list is:\n", vegList)
	fmt.Println("Vegy list is:", len(vegList))
}

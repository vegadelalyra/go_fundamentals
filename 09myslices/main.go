package main

import (
	"fmt"
	"slices"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices land")

	fruitList := []string{"Apple", "Tomato", "Peach"}
	fmt.Printf("Type of fruitlist is %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:])
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)

	fruitList = append(fruitList[:1])
	fmt.Println(fruitList)

	highScores := make([]int, 4)
	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 456
	highScores[3] = 876
	// highScores[4] = 777

	highScores = append(highScores, 555, 666, 321)

	fmt.Printf("highScores: %v\n", highScores)

	fmt.Println("it is sorted?", sort.IntsAreSorted(highScores))
	slices.Sort(highScores)
	fmt.Printf("highScores: %v\n", highScores)
	fmt.Println("it is sorted?", slices.IsSorted(highScores))

	// how to remove a value from slices based on index
	courses := []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Printf("courses: %v\n", courses)
	index := 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Printf("courses: %v\n", courses)
}

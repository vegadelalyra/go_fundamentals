package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in golang")

	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}

	fmt.Printf("days: %v\n", days)

	fmt.Println(">>>>>>> FOR LOOP")
	for day := 0; day < len(days); day++ {
		fmt.Println(days[day])
	}

	fmt.Println(">>>>> FOR RANGE LOOP")
	for i := range days {
		fmt.Println(days[i])
	}

	fmt.Println(">>>>>>>>>>>   FOR x,y RANGE LOOP")
	for index, day := range days {
		fmt.Printf("index is %v and value is %v\n", index, day)
	}

	rogueValue := 1

	for rogueValue < 10 {
		if rogueValue == 2 {
			goto THIS_IS_CRAZY_LABELS
		}

		if rogueValue == 5 {
			rogueValue++
			continue
		}

		fmt.Println("Rogue Value is:", rogueValue)
		rogueValue++
	}

THIS_IS_CRAZY_LABELS:
	fmt.Println("Jumping at PERSEO Realm")
}

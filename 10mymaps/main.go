package main

import "fmt"

func main() {
	fmt.Println("Maps in golang!")

	languages := make(map[string]string)

	languages["JS"] = "JavaScript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python!"

	fmt.Printf("languages: %v\n", languages)
	fmt.Println("JS shorts for", languages["JS"])

	delete(languages, "RB")
	fmt.Printf("languages: %v\n", languages)

	// loops are interesting in golang

	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}
}

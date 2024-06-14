package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://official-joke-api.appspot.com/random_joke?coursename=xd"

func main() {
	fmt.Println("Welcome to handling URLs in golang! :D")
	fmt.Printf("myUrl: %v\n", myUrl)

	// parsing
	result, _ := url.Parse(myUrl)
	fmt.Printf("scheme: %v\n", result.Scheme)
	fmt.Printf("host: %v\n", result.Host)
	fmt.Printf("path: %v\n", result.Path)
	fmt.Printf("port: %v\n", result.Port())
	fmt.Printf("rawQuery: %v\n", result.RawQuery)

	qparams := result.Query() // retrieves raw query values inside url
	fmt.Printf("The type of query params are: %T\n", qparams)

	fmt.Println(qparams["coursename"])

	for _, val := range qparams {
		fmt.Println("Param is", val)
	}

	partsOfUrl := &url.URL{
		Scheme:   "https",
		Host:     "offical-joke-api.appspot.com",
		Path:     "random_joke",
		RawQuery: "coursename=xd",
	}

	anotherURL := partsOfUrl.String()
	fmt.Println(anotherURL)
}

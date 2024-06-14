package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	// go greeter("LOFER")
	greeter("MY")

	websitelist := []string{
		"https://perseo.dev",
		"https://vega.dev",
		"https://otxholock.dev",
		"https://co.future.lag",
		"https://comunism!",
	}

	for _, websitelist := range websitelist {
		getStatusCode(websitelist)
	}
}

func greeter(s string) {
	prints := 0

	for prints < 1 {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
		prints++
	}
}

func getStatusCode(endpoint string) {
	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("Whoopsie in endpoint!")
	} else {
		fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
	}
}

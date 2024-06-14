package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to web verbs!")
	// PerformGetRequest()
	// PerformPostJsonRequest()
	PerformPostFormRequest()
}

const myurl = "http://localhost:8000"

func PerformPostJsonRequest() {
	// fake json payload

	requestBody := strings.NewReader(`
		{
			"coursename":"Let's go with golang",
			"price":0,
			"platform":"lyradePerseo.co"
		}
	`)

	response, err := http.Post(myurl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)

	fmt.Printf("content: %v\n", string(content))
}

func PerformGetRequest() {
	http.Get(myurl)

	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status code: ", response.StatusCode)
	fmt.Println("Content length is: ", response.ContentLength)

	var responseString strings.Builder
	content, _ := io.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)

	fmt.Printf("byteCount: %v\n", byteCount)
	fmt.Printf("responseString: %v\n", responseString.String())

	// fmt.Println(content)
	// fmt.Println(string(content))
}

func PerformPostFormRequest() {
	const urlReq string = myurl + "/form"

	// formdata
	data := url.Values{}
	data.Add("firstname", "perseo")
	data.Add("lastname", "lyra")
	data.Add("email", "go@perseo.vega")

	response, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))
}

package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to JSON in golang")
	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	lyraCourses := []course{
		{
			"ReactJS Bootcamp",
			299,
			"delalyra.perseo.co",
			"abc123",
			[]string{"web-dev", "js"},
		},
		{
			"MERN Bootcamp",
			199,
			"delalyra.perseo.co",
			"hit123",
			[]string{"full-stack", "js"},
		},
		{
			"Angular Bootcamp",
			299,
			"delalyra.perseo.co",
			"bcd123",
			nil,
		},
	}

	// package this data as JSON data
	finalJson, err := json.MarshalIndent(lyraCourses, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	  {
    	"coursename": "ReactJS Bootcamp",
    	"Price": 299,
    	"website": "delalyra.perseo.co",
    	"tags": ["web-dev","js"]
	  }
	`)

	var resiliCourse course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &resiliCourse)
		fmt.Printf("%#v\n", resiliCourse)
	} else {
		fmt.Println("JSON WAS NOT VALID!")
	}

	// some cases where you just want to add data to key value
	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and value is %v and Type is: %T\n", k, v, v)
	}
}

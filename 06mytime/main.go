package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of golang")

	presentTime := time.Now()
	fmt.Printf("presentTime: %v\n", presentTime)
	fmt.Printf("presentTime: %v\n", presentTime.Format("01-02-2013 13:23:32 Monday"))

	createdDate := time.Date(2000, time.April, 20, 23, 23, 0, 0, time.UTC)
	fmt.Printf("createdDate: %v\n", createdDate)
	fmt.Printf("createdDate: %v\n", createdDate.Format("01-02-2000 Monday"))
}

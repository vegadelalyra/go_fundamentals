package main

import "fmt"

var jwtToken int = 300_000 
const LoginToken string = "SDGSGSGD" // Public members are indicated with capital case

func main() {
	var username string = "delalyra"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)


	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat32 float32 = 255.2452534353123412
	fmt.Println(smallFloat32)
	fmt.Printf("Variable is of type: %T \n", smallFloat32)

	var smallFloat64 float64 = 255.2452534353123412
	fmt.Println(smallFloat64)
	fmt.Printf("Variable is of type: %T \n", smallFloat64)

	// default values and some aliases
	fmt.Println("Default types")
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	// implicit type
	fmt.Println("Implicit types")
	var website = "delalyra.cm"
	fmt.Println(website)

	// no var style
	numberOfUser := 300_000.000
	fmt.Printf("numberOfUser: %v\n", numberOfUser)
	
	// global vs local no var style
	fmt.Println(":= style can be used only locally")
	fmt.Println("its global usage is not permitted")
	fmt.Printf("jwtToken: %v\n", jwtToken)

	fmt.Printf("LoginToken: %v\n", LoginToken)
}
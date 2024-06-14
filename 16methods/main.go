package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")

	// no inheritance in golang; No super or parent
	// no classes! there's actually structs, not prototypes neither
	// in go all are types and struct allows us to build an OOP like object

	rlyeh := User{"Rlyeh!", "lyra@gm.co", true, 12}
	fmt.Printf("rlyeh: %v\n", rlyeh)
	fmt.Printf("Rlyeh details are: %+v\n", rlyeh)
	fmt.Printf("Name is %v and email is %v\n", rlyeh.Name, rlyeh.Email)
	rlyeh.GetStatus()
	rlyeh.NewMail()
	fmt.Printf("Name is %v and email is %v\n", rlyeh.Name, rlyeh.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active?", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is:", u.Email)
}

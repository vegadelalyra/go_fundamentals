package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	fmt.Println("Welcome to maths in golang!")

	var myNumOne int = 2
	var myNumTwo float64 = 4

	fmt.Println("The sum is", myNumOne+int(myNumTwo))

	// random number
	// rand.Seed(time.Now().UnixNano()) // before, you had to do this
	// fmt.Println(rand.Intn(5))        // so you could get real rand nums

	// random from crypto
	myRandomCryptoNum, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Printf("myRandomCryptoNum: %v\n", myRandomCryptoNum)
}

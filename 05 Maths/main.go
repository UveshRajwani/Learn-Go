package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	fmt.Println("Welcome to maths in GoLang")
	var num1 = 2
	num2 := 3
	var num3 float64 = 3
	fmt.Println(num1, num2, num3)
	fmt.Printf("type of num1 is:%T\n", num1)
	fmt.Printf("type of num3 is:%T\n", num3)
	fmt.Printf("type of num2 is:%T\n", num2)
	// for rand
	//rand.Seed(time.Now().UnixNano())
	//fmt.Println(rand.Intn(5) + 1)
	// random form crypto

	myrandnum, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(myrandnum)
}

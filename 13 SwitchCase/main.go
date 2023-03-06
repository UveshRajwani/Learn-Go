package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to Switch Case")
	rand.Seed(time.Now().UnixNano())
	dice := rand.Intn(6) + 1
	fmt.Println(dice)

	switch dice {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("IDK")

	}
}

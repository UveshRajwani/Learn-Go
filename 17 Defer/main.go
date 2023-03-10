package main

import "fmt"

func main() {
	defer fmt.Println("This line is defer")
	defer fmt.Println("This line is defer 1")
	defer fmt.Println("This line is defer 2")
	fmt.Println("Welcome to Defer with Golang")
	checkDefer()

}
func checkDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

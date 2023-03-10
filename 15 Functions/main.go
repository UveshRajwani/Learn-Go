package main

import "fmt"

func main() {
	fmt.Println("Welcome to Functions in go")
	result := add(2, 3)
	sayHello()
	fmt.Println(result)
	result2 := proAdder(2, 3, 5, 4, 5, 5, 5, 5, 5, 5, 5)
	fmt.Println(result2)
}
func sayHello() {
	fmt.Println("Hello")
}
func add(e int, b int) int {
	return e + b
}
func proAdder(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}

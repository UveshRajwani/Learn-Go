package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var welcome = "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter your age: ")

	input, _ := reader.ReadString('\n')
	fmt.Println("Your age is: ", input)
	fmt.Printf("Var type: %T \n", input)
}

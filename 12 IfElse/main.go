package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to If and Else")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	loginCount, _ := strconv.ParseFloat(strings.TrimSpace(input), 64)
	var message string
	if loginCount < 10 {
		message = "No Problem"
	}
	if num := loginCount; num < 10 {
		fmt.Println(num)
	}
	fmt.Println(message)
}

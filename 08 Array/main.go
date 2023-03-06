package main

import "fmt"

func main() {
	fmt.Println("Welcome to my array in golang")
	var fruitList [4]string
	fruitList[0] = "apple"
	fruitList[3] = "yes"
	fmt.Println(fruitList)
	fmt.Println(len(fruitList))
	var vegList = [3]string{"d", "w", "d"}
	fmt.Println(vegList)
}

package main

import "fmt"

const UserID = "test2"
const userID2 = "test3"
const UserID3 string = "test3"

func main() {
	var username string = "uvesh"
	fmt.Println(username)
	fmt.Printf("Var type: %T \n", username)
	fmt.Println("----------------------")
	var isWorking bool = false
	fmt.Println(isWorking)
	fmt.Printf("Var type: %T \n", isWorking)
	fmt.Println("----------------------")
	var smallInt uint8 = 255
	fmt.Println(smallInt)
	fmt.Printf("Var type: %T \n", smallInt)
	fmt.Println("----------------------")
	var smallFloat float32 = 255.2222222222
	fmt.Println(smallFloat)
	fmt.Printf("Var type: %T \n", smallFloat)
	fmt.Println("----------------------")
	var bigFloat float64 = 255.2222222222
	fmt.Println(bigFloat)
	fmt.Printf("Var type: %T \n", bigFloat)
	fmt.Println("----------------------")

	var api = "namaz.co.in"
	fmt.Println(api)

	endPoint := "/currentNamaz"
	fmt.Println(endPoint)
	fmt.Println(userID2)
}

package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in go")
	days := []string{"Sunday", "Tuesday", "Friday"}
	fmt.Println(days)
	//for d := 0; d < len(days); d++ {
	//	fmt.Println(days[d])
	//}
	//fmt.Println()
	//for i := range days {
	//	fmt.Println(days[i])
	//}
	//for index, day := range days {
	//	fmt.Printf("the index is %v and value is %v\n", index, day)
	//}
	rValue := 1
	for rValue <= 10 {
		if rValue == 5 {
			break
		}
		fmt.Println(rValue)
		rValue++
	}
}

package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to Slices")
	var fruitList = []int{1, 2, 3}
	fmt.Println(fruitList)
	fruitList = append(fruitList, 5, 4)
	fmt.Println(fruitList)
	sort.Ints(fruitList)
	fmt.Println(sort.IntsAreSorted(fruitList))
	var courses = []string{"js", "python", "html", "css", "c", "go"}
	fmt.Println(courses)
	var index int = 3
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}

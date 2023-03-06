package main

import "fmt"

func main() {
	fmt.Println("Welcome to golang")
	uvesh := User{age: 2, Name: "Uvesh", Email: "test@test.com", Status: false}
	fmt.Println(uvesh.Name)
}

type User struct {
	Name   string
	Email  string
	Status bool
	age    int
}

package main

import "fmt"

func main() {
	fmt.Println("Welcome to golang")
	uvesh := User{Age: 2, Name: "Uvesh", Email: "test@test.com", Status: true}
	uvesh.GetStatus()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is User Active: ", u.Status)
}

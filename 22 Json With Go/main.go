package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string   `json:"coursename"`
	Price    int      `json:"price"`
	Platform string   `json:"platform"`
	Password string   `json:"-"`
	Email    string   `json:"email"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to json with golang")
	EncodeJson()
}
func EncodeJson() {
	lcoCourse := []course{
		{"Python", 0, "Youtube", "test123", "py@go.dev", []string{"py", "fastApi"}},
		{"JavaScript", 0, "Youtube", "test123", "js@go.dev", nil},
	}
	finalJson, err := json.MarshalIndent(lcoCourse, "", "\t")
	checkNilErr(err)
	fmt.Printf("%s\n", finalJson)
}
func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}

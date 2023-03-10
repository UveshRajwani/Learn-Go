package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to files in golang")
	content := "This will go in the file"
	file, err := os.Create("./mylog.txt")
	defer file.Close()
	checkNilErr(err)
	length, err := io.WriteString(file, content)
	checkNilErr(err)
	fmt.Println(length)
	readFile("./mylog.txt")
}
func readFile(filename string) {
	file, err := os.ReadFile(filename)
	checkNilErr(err)
	fmt.Println(string(file))
}
func checkNilErr(err error) {
	if err != nil {
		return
	}
}

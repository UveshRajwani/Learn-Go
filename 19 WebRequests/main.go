package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "http://192.168.1.68"

func main() {
	r, err := http.Get(url)
	checkNilErr(err)
	defer r.Body.Close()
	all, err := io.ReadAll(r.Body)
	checkNilErr(err)
	fmt.Println(string(all))
	fmt.Println(r.Status)
}
func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}

package main

import (
	"fmt"
	"net/url"
)

const myUrl = "http://192.168.1.68:3021/test?noob=1&test=true"

func main() {
	fmt.Println("Welcome to urls in go lang")
	//fmt.Println(myUrl)
	result, _ := url.Parse(myUrl)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)
	fmt.Println(result.Port())
	para := result.Query()
	for key, val := range para {
		fmt.Println("the key is : ", key)
		fmt.Println(val)

	}
	partOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=uvesh",
	}
	fmt.Println(partOfUrl)
}

package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to apis with go lang")
	//PerformGetRequest()
	//PerformPostJsonRequest()
	PerformPostFormDataRequest()
}
func PerformGetRequest() {
	const myUrl = "https://api.namaz.co.in"
	r, err := http.Get(myUrl)
	checkNilErr(err)
	defer r.Body.Close()
	var myData strings.Builder
	data, _ := io.ReadAll(r.Body)
	//fmt.Println(r.StatusCode)
	myData.Write(data)
	fmt.Println(myData.String())
}
func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
func PerformPostJsonRequest() {
	const myUrl = "https://api.namaz.co.in/addUser"
	requestBody := strings.NewReader(`
		{	
		"name":"Uvesh",
		"age":17,
		"City":"Surat"
		}
`)
	post, err := http.Post(myUrl, "application/json", requestBody)
	defer post.Body.Close()
	checkNilErr(err)
	data, _ := io.ReadAll(post.Body)
	fmt.Println(string(data))
}
func PerformPostFormDataRequest() {
	const myUrl = "https://api.namaz.co.in/addUser"
	data := url.Values{}
	data.Add("name", "Uvesh")
	data.Add("age", "17")
	data.Add("eamil", "uvesh@go.dev")
	req, _ := http.PostForm(myUrl, data)
	res, _ := io.ReadAll(req.Body)
	fmt.Println(string(res))
}

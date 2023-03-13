package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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
	DecodeJson()
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
func DecodeJson() {
	const myUrl = "https://api.namaz.co.in"
	myJsonData := []byte(`
        {
                "coursename": "Python",
                "price": 0,
                "platform": "Youtube",
                "email": "py@go.dev",
                "tags": [
                        "py",
                        "fastApi"
                ]
        }`)
	response, _ := http.Get(myUrl)
	data, err := io.ReadAll(response.Body)
	checkNilErr(err)
	defer response.Body.Close()
	checkNilErr(err)
	apiJsonValidCheck := json.Valid(data)
	courseJsonValidCheck := json.Valid(myJsonData)
	if apiJsonValidCheck && courseJsonValidCheck {
		fmt.Println("All Json Is Valid")
		fmt.Println("Api Json Data is \n", string(data))
		var courseJsonData course
		err := json.Unmarshal(myJsonData, &courseJsonData)
		checkNilErr(err)
		fmt.Println("Course Json Data is \n", courseJsonData)
	}

}

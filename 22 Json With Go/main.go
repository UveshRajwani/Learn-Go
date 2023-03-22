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
type currentNamaz struct {
	List []singleMasjid `json:"list"`
	Time timee          `json:"time"`
}
type timee struct {
	Start string `json:"start"`
	End   string `json:"end"`
}
type singleMasjid struct {
	Name   string `json:"name"`
	Img    string `json:"img"`
	Timing map[string]struct {
		JammatTime string `json:"jammat_time"`
	} `json:"timing"`
	//Timing map[string]singleTime `json:"timing"`
}
type singleTime struct {
	JammatTime string `json:"jammat_time"`
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
	const myUrl = "https://api.namaz.co.in/currentnamaz"
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
		var courseJsonData course
		err := json.Unmarshal(myJsonData, &courseJsonData)
		checkNilErr(err)
		//fmt.Println("Course Json Data is \n", courseJsonData)
	}
	//var apiData map[string]interface{}
	var apiData currentNamaz
	err = json.Unmarshal(data, &apiData)
	checkNilErr(err)
	//fmt.Printf("%#v\n", apiData)
	//fmt.Println(apiData)
	for _, val := range apiData.List {
		fmt.Println(val)
	}

}

package main

import "fmt"

func main() {
	fmt.Println("Maps in golang")
	languages := make(map[string]string)
	languages["go"] = "golang"
	languages["py"] = "python"
	languages["js"] = "javascript"
	languages["rb"] = "ruby"
	fmt.Println(languages)
	fmt.Println(languages["js"])
	//delete(languages, "rb")
	//fmt.Println(languages)
	for k, v := range languages {
		fmt.Println(k)
		fmt.Println(v)
	}
}

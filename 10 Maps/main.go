package main

import "fmt"

func main() {
	fmt.Println("Maps in golang")
	languages := make(map[string]string)
	languages["go"] = "golang"
	languages["py"] = "python"
	languages["js"] = "javascript"
	languages["rb"] = "ruby"
	fmt.Println(languages["js"])

}

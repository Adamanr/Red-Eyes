package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string
	Age  int
}

func main() {
	article := User{
		Name: "Artem",
		Age:  12,
	}
	data, err := json.Marshal(article)
	if err != nil {
		fmt.Println("Couldn`t marshal user:", err)
	} else {
		fmt.Println(string(data))
	}
}

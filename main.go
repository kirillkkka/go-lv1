package main

import (
	less7 "github.com/kirillkkka/go-lv1/yamlconf"
	"fmt"
)

func main() {
	_, err := fmt.Println("Hello, world!")
	if err != nil {
		fmt.Println("Error:", err)
	}


conf := less7.Parse()
		fmt.Printf("%v", conf)
}


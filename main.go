package main

import (
	"fmt"
)

func main() {
	_, err := fmt.Println("Hello, world!")
	if err != nil {
		fmt.Println( "Error:", err)
	}
}

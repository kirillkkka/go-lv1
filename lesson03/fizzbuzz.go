package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {

		if i%3 == 0 {
			fmt.Println("Fizz")
		}
		if i%5 == 0 {
			fmt.Println("Buzz")
		}

		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		}

		if i%3 != 0 && i%5 != 0 {
			fmt.Println(i)
		}

		fmt.Println(i)

	}
}

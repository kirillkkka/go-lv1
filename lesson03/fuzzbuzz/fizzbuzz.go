package main

import "fmt"

func main() {
	for i := 0; i < 100; i++ {
		fmt.Println(FizzBuzz(i))
	}
}

func FizzBuzz(n int) string {

	if (n%3 == 0) && (n%5 == 0) {
		return "FizzBuzz"
	} else if n%3 == 0 {
		return "Fizz"
	} else if n%5 == 0 {
		return "Buzz"
	}
	return fmt.Sprintf("%d", n)
}

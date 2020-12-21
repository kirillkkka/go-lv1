package main

import "fmt"

func fib(n int) int {
	cash := make(map[int]int)
	for i := 0; i <= n; i++ {
		var f int
		if i <= 2 {
			f = 1
		} else {
			f = cash[i-1] + cash[i-2]
		}
		cash[i] = f
	}
	fmt.Printf("All line: %d\n", cash) // проверяем мапуБ воводя весь ряд чисел
	return cash[n]
}

func main() {
	var x = fib(9) // указываем значение, для которого ищем число Фибонначи
	fmt.Printf("Fibanocci for your number is: %d\n", x)
}
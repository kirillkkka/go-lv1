package main

import (
	"fmt"
	"math"
	"os"
	//"strconv"
)

func main() {
	var a, b, res float64
	var op string

	fmt.Print("Введите первое число: ")
	fmt.Scanln(&a)

	fmt.Print("Введите второе число: ")
	fmt.Scanln(&b)

	// Вариант с переводом из string во float64 и отработкой ошибок не осилил.
	// Рано полез или так и нужно было делать?
	//
	//var fa, fb float64
	//var err error
	//
	//fa, err := strconv.ParseFloat(a, 64)
	//if err != nil {
	//	fmt.Print(os.Stderr, "Вы ввели не число, перезапустите программу")
	//	os.Exit(1)
	//}
	//
	//fb, err := strconv.ParseFloat(b, 64)
	//if err != nil {
	//	fmt.Print(os.Stderr, "Вы ввели не число, перезапустите программу")
	//	os.Exit(1)
	//}

	fmt.Print("Введите арифметическую операцию (+, -, *, /, ^, %): ")
	fmt.Scanln(&op)

	switch op {
	case "+":
		res = a + b
	case "-":
		res = a - b
	case "*":
		res = a * b
	case "/":
		res = a / b
		if b ==  0 {
			fmt.Fprint(os.Stderr, "Делить на ноль нельзя, перезапустите программу")
			os.Exit(1)
		}
	case "^":
		res = math.Pow(a, b)
	case "%":
		res = math.Mod(a, b)
	default:
		fmt.Println("Операция выбрана неверно")
		os.Exit(1)
	}

	fmt.Printf("Результат выполнения операции: %f\n", res)
}
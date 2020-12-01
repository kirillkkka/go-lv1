package main

import (
	"fmt"
	"math"
	"os"
	"strconv"

	//"strconv"
)

func main() {
	var a, b string
	var res float64
	var fa, fb float64
	var err error
	var op string

	fmt.Print("Введите первое число: ")
	fmt.Scanln(&a)

	fmt.Print("Введите второе число: ")
	fmt.Scanln(&b)

	fa, err = strconv.ParseFloat(a, 64)
	if err != nil {
		fmt.Fprint(os.Stderr, "Вы ввели не число, перезапустите программу\n")
		os.Exit(1)
	}

	fb, err = strconv.ParseFloat(b, 64)
	if err != nil {
		fmt.Fprint(os.Stderr, "Вы ввели не число, перезапустите программу\n")
		os.Exit(1)
	}

	fmt.Print("Введите арифметическую операцию (+, -, *, /, ^, %): ")
	fmt.Scanln(&op)

	switch op {
	case "+":
		res = fa + fb
	case "-":
		res = fa - fb
	case "*":
		res = fa * fb
	case "/":
		res = fa / fb
		if fb ==  0 {
			fmt.Fprint(os.Stderr, "Делить на ноль нельзя, перезапустите программу")
			os.Exit(1)
		}
	case "^":
		res = math.Pow(fa, fb)
	case "%":
		res = math.Mod(fa, fb)
	default:
		fmt.Println("Операция выбрана неверно")
		os.Exit(1)
	}

	fmt.Printf("Результат выполнения операции: %f\n", res)
}
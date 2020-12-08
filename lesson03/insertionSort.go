package main

import (
	"fmt"
)

func insertionSort(sliceToSort []int) []int {
	for i := 1; i < len(sliceToSort); i++ {
		el := sliceToSort[i]
		j := i
		for ; j >= 1 && sliceToSort[j-1] > el; j-- {
			sliceToSort[j] = sliceToSort[j-1]
		}
		sliceToSort[j] = el
	}
	return sliceToSort
}

func main() {
	slice := []int{1, -4, 3, 9, 5, 88, 0, 47}

	fmt.Println("Before insertion sort: ", slice)

	insertionSort(slice)

	fmt.Println("After insertion sort: ", slice)
}
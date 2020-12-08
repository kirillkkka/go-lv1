package main

import (
	"fmt"
)

func bubbleSort(sliceToSort []int) []int {
	if len(sliceToSort) < 2 {
		return sliceToSort
	}
	for i := 0; i < len(sliceToSort); i++ {
		for j := len(sliceToSort) - 1; j >= i+1; j-- {
			if sliceToSort[j] < sliceToSort[j-1] {
				sliceToSort[j], sliceToSort[j-1] = sliceToSort[j-1], sliceToSort[j]
			}
		}
	}
	return sliceToSort
}

func main() {
	slice := []int{1, -4, 3, 9, 5, 88, 0, 47}

	fmt.Println("Before bubble sort: ", slice)

	bubbleSort(slice)

	fmt.Println("After bubble sort: ", slice)
}
package main

import (
	"fmt"
)

func insertionSort(sliceToSort []int) {
	for i := 1; i < len(sliceToSort); i++ {
		x := sliceToSort[i]
		j := i
		for ; j >= 1 && sliceToSort[j-1] > x; j-- {
			sliceToSort[j] = sliceToSort[j-1]
		}
		sliceToSort[j] = x
	}
}

func main() {
	slice := []int{1, -4, 3, 9, 5, 88, 0, 47}

	fmt.Println("Before insertion sort: ", slice)

	insertionSort(slice)

	fmt.Println("After insertion sort: ", slice)
}
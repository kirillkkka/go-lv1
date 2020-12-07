package main

import (
	"fmt"
)

func bubbleSort(sliceToSort []int) {
	size := len(sliceToSort)
	if size < 2 {
		return
	}
	for i := 0; i < size; i++ {
		for j := size - 1; j >= i+1; j-- {
			if sliceToSort[j] < sliceToSort[j-1] {
				sliceToSort[j], sliceToSort[j-1] = sliceToSort[j-1], sliceToSort[j]
			}
		}
	}
}

func main() {
	slice := []int{1, -4, 3, 9, 5, 88, 0, 47}

	fmt.Println("Before bubble sort: ", slice)

	bubbleSort(slice)

	fmt.Println("After bubble sort: ", slice)
}
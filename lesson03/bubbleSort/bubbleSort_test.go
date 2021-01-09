package main

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	nums := []int{1, -4, 3, 9, 5, 88, 0, 47}

	got := BubbleSort(nums)
	want := []int{-4, 0, 1, 3, 5, 9, 47, 88}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v, given: %v", got, want, nums)
	}
}

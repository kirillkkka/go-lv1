package main

import (
	"fmt"
	"testing"
)

var result int
func BenchmarkFib(b *testing.B) {
	var res int
	for i:= 0; i < b.N; i++ {
		res = Fib(20) }
	result = res }


func ExampleFib() {
	n := 5
	fmt.Println(Fib(n))
	// Output: All line: map[0:1 1:1 2:1 3:2 4:3 5:5]
	//5
}
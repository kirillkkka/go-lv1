package main

import "testing"

var result int
func BenchmarkFib(b *testing.B) {
	var res int
	for i:= 0; i < b.N; i++ {
		res = Fib(20) }
	result = res }
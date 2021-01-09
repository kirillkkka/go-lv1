package main

import (
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	cases := []struct {
		in   int
		want string
	}{
		{1, "1"},
		{3, "Fizz"},
		{5, "Buzz"},
		{30, "FizzBuzz"},
	}

	for _, tc := range cases {
		got := FizzBuzz(tc.in)
		if got != tc.want {
			t.Errorf("Fizzbuzz(%d): expected %q, actual %q", tc.in, tc.want, got)
		}
	}

}
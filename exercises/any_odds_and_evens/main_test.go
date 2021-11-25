package main

import "testing"

func TestAnyIsOdd(t *testing.T) {
	for _, tc := range []struct {
		s    []int
		want bool
	}{
		{[]int{}, false},
		{[]int{2}, false},
		{[]int{3}, true},
		{[]int{0, 0, 0}, false},
		{[]int{-2, -4, -8}, false},
		{[]int{-3, -4, -8}, true},
		{[]int{2, 4, 8}, false},
		{[]int{2, 1, 8}, true},
		{[]int{3, 7, 9}, true},
		{[]int{-3, 1, 8}, true},
	} {
		if got := AnyIsOdd(tc.s); got != tc.want {
			t.Errorf("AnyIsOdd(%v) = %v, want %v", tc.s, got, tc.want)
		}
	}
}

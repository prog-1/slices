package main

import "testing"

func TestAnyIsOdd(t *testing.T) {
	for _, tc := range []struct {
		s    []int
		want bool
	}{
		{[]int{}, false},
		{[]int{0}, false},
		{[]int{2}, false},
		{[]int{3}, true},
		{[]int{4, 2}, false},
		{[]int{2, 3}, true},
		{[]int{3, 1, 5}, true},
		{[]int{2, 4, 1}, true},
		{[]int{-1}, true},
	} {
		got := AnyIsOdd(tc.s)
		if got != tc.want {
			t.Errorf("AnyIsOdd(%v) = %v, want = %v", tc.s, got, tc.want)
		}
	}
}
func TestAnyIsEven(t *testing.T) {
	for _, tc := range []struct {
		s    []int
		want bool
	}{
		{[]int{}, false},
		{[]int{0}, true},
		{[]int{1}, false},
		{[]int{2}, true},
		{[]int{2, 8}, true},
		{[]int{3, 5}, false},
		{[]int{2, 7}, true},
		{[]int{3, 1, 5}, false},
		{[]int{2, 4, 6}, true},
		{[]int{-1}, false},
	} {
		got := AnyIsEven(tc.s)
		if got != tc.want {
			t.Errorf("AnyIsEven(%v) = %v, want = %v", tc.s, got, tc.want)
		}
	}
}

package main

import "testing"

func TestAnyIsOdd(t *testing.T) {
	for _, tc := range []struct {
		s    []int
		want bool
	}{
		{[]int{}, false},
		{[]int{2}, false},
		{[]int{1}, true},
		{[]int{2, 3, 4}, true},
		{[]int{8, 6, 4}, false},
		{[]int{5, 1, 6}, true},
		{[]int{-3, 4, 1}, true},
	} {
		got := AnyIsOdd(tc.s)
		if got != tc.want {
			t.Errorf("AnyIsOdd(%v) = %v, want %v", tc.s, got, tc.want)
		}
	}
}
func TestAnyIsEven(t *testing.T) {
	for _, tc := range []struct {
		s    []int
		want bool
	}{
		{[]int{}, false},
		{[]int{2}, true},
		{[]int{1}, false},
		{[]int{4, 4, 5}, true},
		{[]int{3, 1, 5}, false},
		{[]int{6, 8, 12}, true},
		{[]int{-6, 6, 7}, true},
	} {
		got := AnyIsEven(tc.s)
		if got != tc.want {
			t.Errorf("AnyIsEven(%v) = %v, want %v", tc.s, got, tc.want)
		}
	}
}

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
		{[]int{2, 4, 6}, false},
		{[]int{2, 1, 8}, true},
		{[]int{3, 7, 9}, true},
		{[]int{-3, 1, 8}, true},
	} {
		got := AnyIsOdd(tc.s)
		if got != tc.want {
			t.Errorf("AnyIsOdd(%v) - %v, want - %v", tc.s, got, tc.want)
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
		{[]int{2}, true},
		{[]int{3}, false},
		{[]int{2, 4, 6}, true},
		{[]int{2, 1, 8}, true},
		{[]int{3, 7, 9}, false},
		{[]int{-3, 1, 8}, true},
	} {
		got := AnyIsEven(tc.s)
		if got != tc.want {
			t.Errorf("AnyIsEven(%v) - %v, want - %v", tc.s, got, tc.want)
		}
	}
}

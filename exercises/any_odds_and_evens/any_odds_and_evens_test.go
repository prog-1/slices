package main

import (
	"testing"
)

func TestOdd(t *testing.T) {
	for _, tc := range testCasesOdd {
		res := AnyIsOdd(tc.s)
		if res != tc.want {
			t.Errorf("TestOdd(%v) = (%v), want = (%v)",
				tc.s, res, tc.want)
		}
	}
}

func TestEven(t *testing.T) {
	for _, tc := range testCasesEven {
		res := AnyIsEven(tc.s)
		if res != tc.want {
			t.Errorf("TestEven(%v) = (%v), want = (%v)",
				tc.s, res, tc.want)
		}
	}
}

var testCasesEven = []struct {
	s    []int
	want bool
}{
	{[]int{1, 1, 3}, false},
	{[]int{2, 2, 3}, true},
	{[]int{2, 2, 4}, true},
	{[]int{1, 3, 4, 5, 7}, true},
	{[]int{3, 7, 9, 3}, false},
	{[]int{8, 2, 6, 2, 2}, true},
	{[]int{1, 1, 3, 5}, false},
	{[]int{1, 5, 9}, false},
	{[]int{4, 2, 8, 2, 4, 8, 6}, true},
	{[]int{1, 2, 3, 5, 7, 9}, true},
}

var testCasesOdd = []struct {
	s    []int
	want bool
}{
	{[]int{1, 2, 3}, true},
	{[]int{2, 2, 3}, true},
	{[]int{2, 2, 4}, false},
	{[]int{6, 2, 8, 4, 8}, false},
	{[]int{6, 2, 3, 8}, true},
	{[]int{8, 2, 6, 2, 2}, false},
	{[]int{1, 1, 3, 5}, true},
	{[]int{1, 5, 9}, true},
	{[]int{4, 2, 8, 2, 4, 8, 6}, false},
	{[]int{1, 2, 3, 5, 7, 9}, true},
}

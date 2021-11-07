package main

import (
	"testing"
)

func TestOdd(t *testing.T) {
	for _, tc := range testCasesOdd {
		res := AreOddsOnly(tc.s)
		if res != tc.want {
			t.Errorf("TestOdd(%v) = (%v), want = (%v)",
				tc.s, res, tc.want)
		}
	}
}

func TestEven(t *testing.T) {
	for _, tc := range testCasesEven {
		res := AreEvensOnly(tc.s)
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
	{[]int{1, 2, 5, 7}, false},
	{[]int{1, 3, 5, 7}, false},
	{[]int{0, 4, 6, 8}, true},
	{[]int{2, 2, 6, -2}, true},
	{[]int{1, -9, 5, 7}, false},
	{[]int{2, 8, 16, 126}, true},
	{[]int{2, 98, 36, 129}, false},
	{[]int{8, 2, 6, 8}, true},
	{[]int{2, 4, 6, 8}, true},
}

var testCasesOdd = []struct {
	s    []int
	want bool
}{
	{[]int{1, 3, 5}, true},
	{[]int{3, 4, 5}, false},
	{[]int{3, 5, 9}, true},
	{[]int{1, 3, 5, 7, 8}, false},
	{[]int{6, 2, 3, 8}, false},
	{[]int{8, 2, 6, 2, 2}, false},
	{[]int{1, 1, 3, 5}, true},
	{[]int{1, 5, 9}, true},
	{[]int{4, 2, 8, 2, 4, 8, 6}, false},
	{[]int{1, 2, 3, 5, 7, 9}, false},
}

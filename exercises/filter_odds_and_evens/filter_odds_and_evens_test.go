package main

import (
	"reflect"
	"testing"
)

func TestOdd(t *testing.T) {
	for _, tc := range testCasesOdd {
		res := FilterOdds(tc.s)
		if !reflect.DeepEqual(res, tc.want) {
			t.Errorf("TestOdd(%v) = (%v), want = (%v)",
				tc.s, res, tc.want)
		}
	}
}

func TestEven(t *testing.T) {
	for _, tc := range testCasesEven {
		res := FilterEvens(tc.s)

		if !reflect.DeepEqual(res, tc.want) {
			t.Errorf("TestEven(%v) = (%v), want = (%v)",
				tc.s, res, tc.want)
		}
	}
}

var testCasesEven = []struct {
	s    []int
	want []int
}{
	{[]int{1, 2, 5, 7}, []int{2}},
	{[]int{1, 3, 5, 7}, []int{}},
	{[]int{0, 4, 6, 8}, []int{0, 4, 6, 8}},
	{[]int{2, 2, 6, -2}, []int{2, 2, 6, -2}},
	{[]int{1, -9, 6, 8}, []int{6, 8}},
	{[]int{2, 8, 16, 126}, []int{2, 8, 16, 126}},
	{[]int{2, 98, 36, 129}, []int{2, 98, 36}},
	{[]int{8, 2, 6, 8}, []int{8, 2, 6, 8}},
	{[]int{2, 4, 6, 8}, []int{2, 4, 6, 8}},
}

var testCasesOdd = []struct {
	s    []int
	want []int
}{
	{[]int{1, 2, 5, 7}, []int{1, 5, 7}},
	{[]int{1, 3, 5, 7}, []int{1, 3, 5, 7}},
	{[]int{0, 4, 6, 8}, []int{}},
	{[]int{2, 2, 6, -2}, []int{}},
	{[]int{1, -9, 6, 8}, []int{1, -9}},
	{[]int{2, 8, 16, 126}, []int{}},
	{[]int{2, 98, 36, 129}, []int{129}},
	{[]int{8, 2, 6, 8}, []int{}},
	{[]int{2, 4, 6, 8}, []int{}},
}

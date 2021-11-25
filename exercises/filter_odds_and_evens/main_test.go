package main

import "testing"

func equal(x, y []int) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}
func TestFilterEvens(t *testing.T) {
	for _, tc := range []struct {
		s    []int
		want []int
	}{
		{[]int{}, []int{}},
		{[]int{0}, []int{0}},
		{[]int{0, 1}, []int{0}},
		{[]int{-1, -2}, []int{-2}},
		{[]int{3, 7, 4, 9}, []int{4}},
		{[]int{4, 6, 8, 1, 3, 91}, []int{4, 6, 8}},
	} {

		if got := FilterEvens(tc.s); !equal(got, tc.want) {
			t.Errorf("FilterEvens(%v) = %v, want = %v", tc.s, got, tc.want)
		}
	}
}

func TestFilterOdds(t *testing.T) {
	for _, tc := range []struct {
		s    []int
		want []int
	}{
		{[]int{}, []int{}},
		{[]int{0}, []int{}},
		{[]int{0, 1}, []int{1}},
		{[]int{-1, -2}, []int{-1}},
		{[]int{3, 7, 4, 9}, []int{3, 7, 9}},
		{[]int{4, 6, 8, 1, 3, 91}, []int{1, 3, 91}},
		{[]int{-4, 6, -8, 1, -3, -91}, []int{1, -3, -91}},
	} {
		got := FilterOdds(tc.s)
		if !equal(got, tc.want) {
			t.Errorf("FilterOdds(%v) = %v, want = %v", tc.s, got, tc.want)
		}
	}
}

package main

import "testing"

func TestFilterOdds(t *testing.T) {
	for _, tc := range []struct {
		s    []int
		want []int
	}{
		{[]int{2, 5, 6, 7}, []int{5, 7}},
		{[]int{}, []int{}},
		{[]int{2, 4, 6}, []int{}},
		{[]int{3, 7, 1}, []int{3, 7, 1}},
		{[]int{-2, 5, 4}, []int{5}},
		{[]int{-3, 1, 5}, []int{-3, 1, 5}},
		{[]int{1}, []int{1}},
		{[]int{2}, []int{}},
	} {
		got := FilterOdds(tc.s)
		if !equal(got, tc.want) {
			t.Errorf("FilterOdds(%v) = %v, want = %v", tc.s, got, tc.want)
		}

	}
}

func TestFilterEvens(t *testing.T) {
	for _, tc := range []struct {
		s    []int
		want []int
	}{
		{[]int{}, []int{}},
		{[]int{1}, []int{}},
		{[]int{2}, []int{2}},
		{[]int{2, 4, 6}, []int{2, 4, 6}},
		{[]int{3, 9, 7}, []int{}},
		{[]int{4, 5, 8}, []int{4, 8}},
		{[]int{-3, -2, 5}, []int{-2}},
	} {
		got := FilterEvens(tc.s)
		if !equal(got, tc.want) {
			t.Errorf("FilterEvens(%v) = %v, want = %v", tc.s, got, tc.want)
		}
	}
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

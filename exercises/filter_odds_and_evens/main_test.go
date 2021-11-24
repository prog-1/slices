package main

import "testing"

func TestFilterOdds(t *testing.T) {
	for _, tc := range []struct {
		s    []int
		want []int
	}{
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{2}, []int{}},
		{[]int{5, 2, 6}, []int{5}},
		{[]int{3, 9, 5}, []int{3, 9, 5}},
		{[]int{-2, 8, -6}, []int{}},
		{[]int{-3, 3, 3}, []int{-3, 3, 3}},
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
		{[]int{7, 2, 1}, []int{2}},
		{[]int{2, 6, 4}, []int{2, 6, 4}},
		{[]int{-5, 7, 3}, []int{}},
		{[]int{-6, 8, 2}, []int{-6, 8, 2}},
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

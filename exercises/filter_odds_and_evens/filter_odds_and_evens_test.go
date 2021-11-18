package main

import "testing"

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

func TestFilterOdds(t *testing.T) {
	for _, tc := range []struct {
		s    []int
		want []int
	}{
		{[]int{}, []int{}},
		{[]int{0}, []int{}},
		{[]int{2}, []int{}},
		{[]int{3}, []int{3}},
		{[]int{2, 4}, []int{}},
		{[]int{2, 1}, []int{1}},
		{[]int{3, 7, 9}, []int{3, 7, 9}},
		{[]int{2, 3, 6}, []int{3}},
		{[]int{-3}, []int{-3}},
		{[]int{-3, -7}, []int{-3, -7}},
		{[]int{-3, -2}, []int{-3}},
	} {
		got := make([]int, len(tc.s))
		copy(got, tc.s)
		got = FilterOdds(got)
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
		{[]int{0}, []int{0}},
		{[]int{2}, []int{2}},
		{[]int{3}, []int{}},
		{[]int{2, 4}, []int{2, 4}},
		{[]int{2, 1}, []int{2}},
		{[]int{3, 7, 9}, []int{}},
		{[]int{2, 3, 6}, []int{2, 6}},
		{[]int{-3}, []int{}},
		{[]int{-2, -6}, []int{-2, -6}},
		{[]int{-3, -2}, []int{-2}},
	} {
		got := make([]int, len(tc.s))
		copy(got, tc.s)
		got = FilterEvens(got)
		if !equal(got, tc.want) {
			t.Errorf("FilterEvens(%v) = %v, want = %v", tc.s, got, tc.want)
		}
	}
}

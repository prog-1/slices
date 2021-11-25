package main

import "testing"

func TestFilterOdds(t *testing.T) {
	for _, tc := range []struct {
		s    []int
		want []int
	}{
		{[]int{}, []int{}},
		{[]int{0}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{2}, []int{}},
		{[]int{3}, []int{3}},
		{[]int{1, 3}, []int{1, 3}},
		{[]int{6, 10}, []int{}},
		{[]int{2, 5}, []int{5}},
		{[]int{3, 4, 9}, []int{3, 9}},
		{[]int{-1, 6, -7}, []int{-1, -7}},
		{[]int{-4, 0, 8}, []int{}},
	} {
		got := FilterOdds(tc.s)
		if !equal(got, tc.want) {
			t.Errorf("FilterOdds(%v) - %v, want - %v", tc.s, got, tc.want)
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
		{[]int{1}, []int{}},
		{[]int{2}, []int{2}},
		{[]int{3}, []int{}},
		{[]int{1, 3}, []int{}},
		{[]int{6, 10}, []int{6, 10}},
		{[]int{2, 5}, []int{2}},
		{[]int{3, 4, 9}, []int{4}},
		{[]int{-1, 6, -7}, []int{6}},
		{[]int{-4, 0, 8}, []int{-4, 0, 8}},
	} {
		got := FilterEvens(tc.s)
		if !equal(got, tc.want) {
			t.Errorf("FilterEvens(%v) - %v, want - %v", tc.s, got, tc.want)
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

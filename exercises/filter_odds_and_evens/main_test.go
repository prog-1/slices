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
		{[]int{2}, []int{}},
		{[]int{3}, []int{3}},
		{[]int{2, 4, 6}, []int{}},
		{[]int{2, 1, 8}, []int{1}},
		{[]int{3, 7, 9}, []int{3, 7, 9}},
		{[]int{-3, 1, 8}, []int{-3, 1}},
	} {
		got := FilterOdds(tc.s)
		if !equal(got, tc.want) {
			t.Errorf("Sort(%v) = %v, want = %v", tc.s, got, tc.want)
		}
	}
}
func TestFilterEvens(t *testing.T) {
	for _, tc := range []struct {
		s    []int
		want []int
	}{
		{[]int{}, []int{}},
		{[]int{2}, []int{2}},
		{[]int{3}, []int{}},
		{[]int{2, 4, 6}, []int{2, 4, 6}},
		{[]int{2, 1, 8}, []int{2, 8}},
		{[]int{3, 7, 9}, []int{}},
		{[]int{-3, 1, 8}, []int{8}},
	} {
		got := FilterEvens(tc.s)
		if !equal(got, tc.want) {
			t.Errorf("FilterEvens(%v) = %v, want = %v", tc.s, got, tc.want)
		}
	}

}

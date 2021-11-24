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
		{[]int{3}, []int{3}},
		{[]int{2}, []int{}},
		{[]int{3, 5, 7}, []int{3, 5, 7}},
		{[]int{3, 2, 5}, []int{3, 5}},
		{[]int{-1, 3, -2}, []int{-1, 3}},
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
		{[]int{2}, []int{2}},
		{[]int{3}, []int{}},
		{[]int{4, 4, 4}, []int{4, 4, 4}},
		{[]int{2, -1, 8}, []int{2, 8}},
		{[]int{7, 7, 7}, []int{}},
	} {
		got := FilterEvens(tc.s)
		if !equal(got, tc.want) {
			t.Errorf("FilterEvens(%v) = %v, want = %v", tc.s, got, tc.want)
		}
	}

}

package main

import "testing"

func TestSort1(t *testing.T) {
	for _, tc := range []struct {
		s    []int
		want []int
	}{
		{},
		{[]int{}, []int{}},
		{[]int{0}, []int{0}},
		{[]int{1}, []int{1}},
		{[]int{2, 1}, []int{1, 2}},
		{[]int{3, 9, 1}, []int{1, 3, 9}},
		{[]int{2, 93, 16, 476, 9}, []int{2, 9, 16, 93, 476}},
		{[]int{-1, -2}, []int{-2, -1}},
		{[]int{2, 9, -1}, []int{-1, 2, 9}},
	} {
		got := make([]int, len(tc.s))
		copy(got, tc.s)
		Sort1(got)
		if !equal(got, tc.want) {
			t.Errorf("Sort(%v) = %v, want = %v", tc.s, got, tc.want)
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

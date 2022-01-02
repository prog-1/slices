package main

import "testing"

func TestSort1(t *testing.T) {
	for _, tc := range []struct {
		s    []int
		want []int
	}{
		{[]int{1, 2, 3}, []int{1, 2, 3}},
		{[]int{1, 3, 2}, []int{1, 2, 3}},
		{[]int{2, 1, 3}, []int{1, 2, 3}},
		{[]int{2, 3, 1}, []int{1, 2, 3}},
		{[]int{3, 1, 2}, []int{1, 2, 3}},
		{[]int{3, -2, 3}, []int{-2, 3, 3}},
		{[]int{-2, -2, -1}, []int{-2, -2, -1}},
	} {
		got := Sort1(tc.s)
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

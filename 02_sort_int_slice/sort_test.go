package main

import "testing"

func TestSort(t *testing.T) {
	for _, tc := range []struct {
		s    []int
		want []int
	}{
		{},                 // Testing zero values.
		{nil, nil},         // The same as before.
		{[]int{}, []int{}}, // Empty slice isn't the same.
		{[]int{1}, []int{1}},
		{[]int{1, 1}, []int{1, 1}},
		{[]int{-1, 1}, []int{-1, 1}},
		{[]int{1, -1}, []int{-1, 1}},
		{[]int{3, 2, 1, 0, -1, -2, -3}, []int{-3, -2, -1, 0, 1, 2, 3}},
	} {
		// We need to copy tc.s, since we sort in-place.
		got := make([]int, len(tc.s))
		copy(got, tc.s)
		Sort(got)
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

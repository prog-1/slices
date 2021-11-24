package main

import "testing"

func TestSort1(t *testing.T) {
	for _, tc := range []struct {
		s    []int
		want []int
	}{
		{[]int{}, []int{}},
		{[]int{1, 0}, []int{0, 1}},
		{[]int{0, 1}, []int{0, 1}},
		{[]int{1}, []int{1}},
		{[]int{1, 1}, []int{1, 1}},
		{[]int{-1, 1}, []int{-1, 1}},
		{[]int{1, -1}, []int{-1, 1}},
		{[]int{7, 5, 2}, []int{2, 5, 7}},
		{[]int{-3, 8, 1}, []int{-3, 1, 8}},
	} {
		got := make([]int, len(tc.s))
		copy(got, tc.s)
		Sort1(got)
		if !equal(got, tc.want) {
			t.Errorf("Sort(%v) = %v, want = %v", tc.s, got, tc.want)
		}
	}
}

func TestSort2(t *testing.T) {
	for _, tc := range []struct {
		s    []int
		want []int
	}{
		{[]int{}, []int{}},
		{[]int{1, 0}, []int{0, 1}},
		{[]int{0, 1}, []int{0, 1}},
		{[]int{1}, []int{1}},
		{[]int{1, 1}, []int{1, 1}},
		{[]int{-1, 1}, []int{-1, 1}},
		{[]int{1, -1}, []int{-1, 1}},
		{[]int{7, 5, 2}, []int{2, 5, 7}},
		{[]int{-3, 8, 1}, []int{-3, 1, 8}},
	} {
		got := make([]int, len(tc.s))
		copy(got, tc.s)
		Sort2(got)
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

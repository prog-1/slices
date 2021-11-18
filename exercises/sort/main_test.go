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

func TestSort1(t *testing.T) {
	for _, tc := range []struct {
		s    []int
		want []int
	}{
		{[]int{}, []int{}},
		{[]int{0}, []int{0}},
		{[]int{1}, []int{1}},
		{[]int{-2, 1, -5}, []int{-5, -2, 1}},
		{[]int{4, 98, 6}, []int{4, 6, 98}},
		{[]int{1, 2, 3}, []int{1, 2, 3}},
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
		{[]int{0}, []int{0}},
		{[]int{1}, []int{1}},
		{[]int{-2, 1, -5}, []int{-5, -2, 1}},
		{[]int{4, 98, 6}, []int{4, 6, 98}},
		{[]int{1, 2, 3}, []int{1, 2, 3}},
	} {
		got := make([]int, len(tc.s))
		copy(got, tc.s)
		Sort2(got)
		if !equal(got, tc.want) {
			t.Errorf("Sort(%v) = %v, want = %v", tc.s, got, tc.want)
		}
	}
}

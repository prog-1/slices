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
		{[]int{4, 2}, []int{2, 4}},
		{[]int{2, 1}, []int{1, 2}},
		{[]int{5, 7, 3}, []int{3, 5, 7}},
		{[]int{2, 6, 3}, []int{2, 3, 6}},
		{[]int{-3, 5, 2}, []int{-3, 2, 5}},
		{[]int{-3, -7, 4}, []int{-7, -3, 4}},
	} {
		got := make([]int, len(tc.s))
		copy(got, tc.s)
		got = Sort1(got)
		if !equal(got, tc.want) {
			t.Errorf("Sort1(%v) = %v, want = %v", tc.s, got, tc.want)
		}
	}
}

func TestSort2(t *testing.T) {
	for _, tc := range []struct {
		s    []int
		want []int
	}{
		{[]int{}, []int{}},
		{[]int{4, 2}, []int{2, 4}},
		{[]int{2, 1}, []int{1, 2}},
		{[]int{5, 7, 3}, []int{3, 5, 7}},
		{[]int{2, 6, 3}, []int{2, 3, 6}},
		{[]int{-3, 5, 2}, []int{-3, 2, 5}},
		{[]int{-3, -7, 4}, []int{-7, -3, 4}},
	} {
		got := make([]int, len(tc.s))
		copy(got, tc.s)
		got = Sort2(got)
		if !equal(got, tc.want) {
			t.Errorf("Sort2(%v) = %v, want = %v", tc.s, got, tc.want)
		}
	}
}

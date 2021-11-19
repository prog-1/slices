package main

import "testing"

func TestSort1(t *testing.T) {
	for _, tc := range []struct {
		s    []int
		want []int
	}{
		{[]int{}, []int{}},
		{[]int{0}, []int{0}},
		{[]int{1}, []int{1}},
		{[]int{3, 5}, []int{3, 5}},
		{[]int{8, 7}, []int{7, 8}},
		{[]int{2, 4, 9}, []int{2, 4, 9}},
		{[]int{8, 6, 7}, []int{6, 7, 8}},
		{[]int{1, 3, -2}, []int{-2, 1, 3}},
		{[]int{0, -6, 1}, []int{-6, 0, 1}},
	} {
		got := Sort1(tc.s)
		if !equal(got, tc.want) {
			t.Errorf("Sort1(%v) - %v, want - %v", tc.s, got, tc.want)
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
		{[]int{3, 5}, []int{3, 5}},
		{[]int{8, 7}, []int{7, 8}},
		{[]int{2, 4, 9}, []int{2, 4, 9}},
		{[]int{8, 6, 7}, []int{6, 7, 8}},
		{[]int{1, 3, -2}, []int{-2, 1, 3}},
		{[]int{0, -6, 1}, []int{-6, 0, 1}},
	} {
		got := Sort2(tc.s)
		if !equal(got, tc.want) {
			t.Errorf("Sort2(%v) - %v, want - %v", tc.s, got, tc.want)
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

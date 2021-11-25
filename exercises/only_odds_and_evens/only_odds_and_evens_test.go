package main

import "testing"

func TestAreOddsOnly(t *testing.T) {
	for _, tc := range []struct {
		s    []int
		want bool
	}{
		{[]int{0}, false},
		{[]int{2}, false},
		{[]int{3}, true},
		{[]int{2, 4, 6}, false},
		{[]int{2, 1, 8}, false},
		{[]int{3, 7, 9}, true},
		{[]int{-1, 5, -13}, true},
		{[]int{-4, 8, -6}, false},
		{[]int{-3, 1, 10}, false},
	} {
		got := AreOddsOnly(tc.s)
		if got != tc.want {
			t.Errorf("AreOddsOnly(%v) - %v, want - %v", tc.s, got, tc.want)
		}
	}
}

func TestAreEvenOnly(t *testing.T) {
	for _, tc := range []struct {
		s    []int
		want bool
	}{
		{[]int{0}, true},
		{[]int{2}, true},
		{[]int{3}, false},
		{[]int{2, 4, 6}, true},
		{[]int{2, 1, 8}, false},
		{[]int{3, 7, 9}, false},
		{[]int{-1, 5, -13}, false},
		{[]int{-4, 8, -6}, true},
		{[]int{-3, 1, 10}, false},
	} {
		got := AreEvenOnly(tc.s)
		if got != tc.want {
			t.Errorf("AreEvenOnly(%v) - %v, want - %v", tc.s, got, tc.want)
		}
	}
}

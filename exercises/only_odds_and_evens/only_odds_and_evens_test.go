package main

import "testing"

func TestAreOddsOnly(t *testing.T) {
	for _, tc := range []struct {
		s    []int
		want bool
	}{
		{[]int{}, true},
		{[]int{0}, false},
		{[]int{1}, true},
		{[]int{2}, false},
		{[]int{2, 4, 5}, false},
		{[]int{1, 5, 7}, true},
		{[]int{2, 8, 6}, false},
		{[]int{-3, -1, 3}, true},
		{[]int{-2, 4, 6}, false},
		{[]int{-1, 5, 6}, false},
	} {
		got := AreOddsOnly(tc.s)
		if got != tc.want {
			t.Errorf("AreOddsOnly(%v) = %v, want %v", tc.s, got, tc.want)
		}
	}
}

func TestAreEvenOnly(t *testing.T) {
	for _, tc := range []struct {
		s    []int
		want bool
	}{
		{[]int{}, true},
		{[]int{0}, true},
		{[]int{1}, false},
		{[]int{2}, true},
		{[]int{4, 5}, false},
		{[]int{2, 6, 8}, true},
		{[]int{5, 8, 7}, false},
		{[]int{-1, 2, 3}, false},
		{[]int{-6, -4, -2}, true},
	} {
		got := AreEvenOnly(tc.s)
		if got != tc.want {
			t.Errorf("AreEvenOnly(%v) = %v , want = %v", tc.s, got, tc.want)
		}
	}
}

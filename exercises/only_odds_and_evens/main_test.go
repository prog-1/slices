package main

import "testing"

func TestAreOddsOnly(t *testing.T) {
	for _, tc := range []struct {
		s    []int
		want bool
	}{
		{[]int{}, true},
		{[]int{1}, true},
		{[]int{2}, false},
		{[]int{4, 5, 6}, false},
		{[]int{7, 9, 5}, true},
		{[]int{-2, 5, 7}, false},
		{[]int{-9, 7, 1}, true},
	} {
		got := AreOddsOnly(tc.s)
		if got != tc.want {
			t.Errorf("AreOddsOnly(%v) = %v , want = %v", tc.s, got, tc.want)
		}
	}
}
func TestAreEvenOnly(t *testing.T) {
	for _, tc := range []struct {
		s    []int
		want bool
	}{
		{[]int{}, true},
		{[]int{1}, false},
		{[]int{2}, true},
		{[]int{4, 5, 6}, false},
		{[]int{4, 8, 2}, true},
		{[]int{-1, 5, 6}, false},
		{[]int{-4, 6, 6}, true},
	} {
		got := AreEvenOnly(tc.s)
		if got != tc.want {
			t.Errorf("AreEvenOnly(%v) = %v , want = %v", tc.s, got, tc.want)
		}
	}
}

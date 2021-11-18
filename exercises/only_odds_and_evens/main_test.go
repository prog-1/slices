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
		{[]int{2, 3}, false},
		{[]int{2, 2, 4}, false},
		{[]int{1, 3, 5, 6}, false},
		{[]int{-1, -3, -5}, true},
		{[]int{1, 21, 1286731}, true},
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
		{[]int{0}, true},
		{[]int{1}, false},
		{[]int{2}, true},
		{[]int{2, 3}, false},
		{[]int{2, 2, 4}, true},
		{[]int{1, 3, 5, 9, 6}, false},
		{[]int{-2, -6, 8}, true},
		{[]int{1, 76, 21}, false},
		{[]int{12122, 9182, 854}, true},
	} {
		got := AreEvenOnly(tc.s)
		if got != tc.want {
			t.Errorf("AreEvenOnly(%v) = %v , want = %v", tc.s, got, tc.want)
		}
	}
}

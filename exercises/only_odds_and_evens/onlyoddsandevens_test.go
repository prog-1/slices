package main

import "testing"

func TestAreEvenOnly(t *testing.T) {
	for _, tc := range []struct {
		s    []int
		want bool
	}{
		{[]int{}, true},
		{[]int{0}, false},
		{[]int{1}, true},
		{[]int{2}, false},
		{[]int{1, 5, 7}, true},
		{[]int{2, 8, 6}, false},
		{[]int{2, 4, 5}, false},
	} {
		got := AreEvenOnly(tc.s)
		if got != tc.want {
			t.Errorf("AreEvenOnly(%v) = %v , want = %v", tc.s, got, tc.want)
		}
	}
}

func TestAreOddsOnly(t *testing.T) {
	for _, tc := range []struct {
		s    []int
		want bool
	}{
		{[]int{}, true},
		{[]int{1}, true},
		{[]int{2}, false},
	} {
		got := AreOddsOnly(tc.s)
		if got != tc.want {
			t.Errorf("AreOddsOnly(%v) = %v, want %v", tc.s, got, tc.want)
		}

	}
}

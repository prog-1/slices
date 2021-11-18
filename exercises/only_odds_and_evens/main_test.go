package main

import "testing"

func TestAreOddsOnly(t *testing.T) {
	for _, tc := range []struct {
		s    []int
		want bool
	}{
		{[]int{}, true},
		{[]int{2}, false},
		{[]int{3}, true},
		{[]int{2, 4, 6}, false},
		{[]int{2, 1, 8}, false},
		{[]int{3, 7, 9}, true},
		{[]int{-3, 1, 8}, false},
	} {
		got := AreOddsOnly(tc.s)
		if got != tc.want {
			t.Errorf("AreOddsOnly(%v) = %v, want = %v", tc.s, got, tc.want)
		}
	}
}
func TestAreEvenOnly(t *testing.T) {
	for _, tc := range []struct {
		s    []int
		want bool
	}{
		{[]int{}, true},
		{[]int{2}, true},
		{[]int{3}, false},
		{[]int{2, 4, 6}, true},
		{[]int{2, 1, 8}, false},
		{[]int{3, 7, 9}, false},
		{[]int{-3, 1, 8}, false},
	} {
		got := AreEvenOnly(tc.s)
		if got != tc.want {
			t.Errorf("AnyIsOdd(%v) = %v, want = %v", tc.s, got, tc.want)
		}
	}

}

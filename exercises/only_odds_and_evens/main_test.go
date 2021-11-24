package main

import "testing"

func TestAreOddsOnly(t *testing.T) {
	for _, tc := range []struct {
		s    []int
		want bool
	}{
		{[]int{-101}, true},
		{[]int{2}, false},
		{[]int{3}, true},
		{[]int{2, 7}, false},
		{[]int{1, 121}, true},
		{[]int{2, 4, 8}, false},
		{[]int{8, 6, 2}, false},
		{[]int{44, 44, 100}, false},
		{[]int{2, 600, 600}, false},
		{[]int{10, 10, 10}, false},
		{[]int{2, 100, 3}, false},
		{[]int{4, 4, 1}, false},
		{[]int{133, 200, 200}, false},
		{[]int{1, 13, 21}, true},
		{[]int{345, 21, 1}, true},
		{[]int{13, 21, 9}, true},
		{[]int{2, 1, 8}, false},
		{[]int{3, 7, 9}, true},
		{[]int{-3, 1, 9}, true},
		{[]int{-200, -7, -1}, false},
		{[]int{-1, -3, 13}, true},
		{[]int{2, -7, -7}, false},
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
		{[]int{2}, true},
		{[]int{-10}, true},
		{[]int{88}, true},
		{[]int{13877}, false},
		{[]int{56, 2, 8}, true},
		{[]int{56, 56, 28}, true},
		{[]int{2, 3, 3}, false},
		{[]int{100, 100, 100}, true},
		{[]int{13, 13, 13}, false},
		{[]int{2, 8, 10}, true},
		{[]int{10, 8, 4}, true},
		{[]int{7, 9, 15}, false},
		{[]int{25, 3, 10001}, false},
		{[]int{48, 100, 9}, false},
		{[]int{33, 11, 2}, false},
		{[]int{4, 9, 3}, false},
		{[]int{3, 7, 2}, false},
		{[]int{-2, 90, 8}, true},
		{[]int{-1, -3, -5}, false},
		{[]int{-2, -4, 6}, true},
		{[]int{-1, -5, 699}, false},
	} {
		got := AreEvenOnly(tc.s)
		if got != tc.want {
			t.Errorf("AreEvenOnly(%v) = %v, want = %v", tc.s, got, tc.want)
		}
	}
}

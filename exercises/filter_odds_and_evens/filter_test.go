package main

import "testing"

func TestFilterOdds(t *testing.T) {
	for _, tc := range []struct {
		s    []int
		want []int
	}{
		{[]int{-101}, []int{-101}},
		{[]int{3}, []int{3}},
		{[]int{1, 121}, []int{1, 121}},
		{[]int{2, 4, 8}, []int{}},
		{[]int{8, 6, 2}, []int{}},
		{[]int{44, 44, 100, 200, 13370}, []int{}},
		{[]int{2, 600, 600}, []int{}},
		{[]int{10, 10, 10}, []int{}},
		{[]int{2, 100, 3}, []int{3}},
		{[]int{4, 4, 1}, []int{1}},
		{[]int{133, 200, 200}, []int{133}},
		{[]int{1, 13, 21}, []int{1, 13, 21}},
		{[]int{345, 21, 1, 7, 1457, 8889}, []int{345, 21, 1, 7, 1457, 8889}},
		{[]int{13, 21, 9}, []int{13, 21, 9}},
		{[]int{2, 1, 8}, []int{1}},
		{[]int{3, 7, 9, 6, 8}, []int{3, 7, 9}},
		{[]int{-3, 1, 9}, []int{-3, 1, 9}},
		{[]int{-200, -7, -1}, []int{-7, -1}},
		{[]int{-1, -3, 48, 1002, 13}, []int{-1, -3, 13}},
		{[]int{2, -7, -7}, []int{-7, -7}},
	} {
		got := FilterOdds(tc.s)
		if !equal(got, tc.want) {
			t.Errorf("FilterOdds(%v) = %v, want = %v", tc.s, got, tc.want)
		}
	}
}

func TestFilterEvens(t *testing.T) {
	for _, tc := range []struct {
		s    []int
		want []int
	}{
		{[]int{-10}, []int{-10}},
		{[]int{13877}, []int{}},
		{[]int{56, 2, 8, 100, 300, 17390}, []int{56, 2, 8, 100, 300, 17390}},
		{[]int{56, 7, 7, 31, 56, 28}, []int{56, 56, 28}},
		{[]int{2, 3, 3}, []int{2}},
		{[]int{100, 100, 100}, []int{100, 100, 100}},
		{[]int{13, 13, 13, 1}, []int{}},
		{[]int{2, 8, 10}, []int{2, 8, 10}},
		{[]int{10, 8, 4}, []int{10, 8, 4}},
		{[]int{7, 9, 8, -99, 15}, []int{8}},
		{[]int{25, 3, 10001}, []int{}},
		{[]int{48, 100, 9}, []int{48, 100}},
		{[]int{33, 11, 2}, []int{2}},
		{[]int{4, 9, 3}, []int{4}},
		{[]int{-2, 90, 8, -100, 720}, []int{-2, 90, 8, -100, 720}},
		{[]int{-1, -3, -5}, []int{}},
		{[]int{-2, -4, 77, 33}, []int{-2, -4}},
		{[]int{-1, -5, 699}, []int{}},
	} {
		got := FilterEvens(tc.s)
		if !equal(got, tc.want) {
			t.Errorf("FilterEvens(%v) = %v, want = %v", tc.s, got, tc.want)
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

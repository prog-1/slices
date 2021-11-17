package main

import "testing"

func TestOnlyOdd(t *testing.T) {
	for _, tc := range []struct {
		s    []int
		want []int
	}{
		{[]int{}, []int{}},
		{[]int{1, 3, 4}, []int{1, 3}},
		{[]int{1, 3, 5}, []int{1, 3, 5}},
		{[]int{2, 4, 6}, []int{}},
		{[]int{0}, []int{}},
		{[]int{1, 9, 3}, []int{1, 9, 3}},
	} {
		got := OnlyOdd(tc.s)
		if got != tc.want {
			t.Errorf("AllOdd %v = %v, want %v", tc.s, got, tc.want)
		}
	}
}
func TestOnlyEven(t *testing.T) {
	for _, tc := range []struct {
		s    []int
		want []int
	}{
		{[]int{}, []int{}},
		{[]int{1, 3, 4}, []int{4}},
		{[]int{1, 3, 5}, []int{}},
		{[]int{2, 4, 6}, []int{2, 4, 6}},
		{[]int{2, 4, 7}, []int{2, 4}},
		{[]int{0}, []int{}},
		{[]int{1, 9, 3}, []int{}},
	} {
		got := OnlyEven(tc.s)
		if got != tc.want {
			t.Errorf("AllEven %v = %v, want %v", tc.s, got, tc.want)
		}
	}
}

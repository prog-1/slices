package main

import (
	"reflect"
	"testing"
)

func equal(a, b [][]int) bool {
	return reflect.DeepEqual(a, b)
}
func TestSpiral(t *testing.T) {
	for _, tc := range []struct {
		x, y int
		want [][]int
	}{
		{0, 0, [][]int{}},
		{1, 1, [][]int{
			{1},
		}},
		{7, 5, [][]int{
			{1, 2, 3, 4, 5, 6, 7},
			{20, 21, 22, 23, 24, 25, 8},
			{19, 32, 33, 34, 35, 26, 9},
			{18, 31, 30, 29, 28, 27, 10},
			{17, 16, 15, 14, 13, 12, 11},
		}},
	} {
		got := Spiral(tc.x, tc.y)
		if !equal(got, tc.want) {
			t.Errorf("Spiral(%v, %v) = %v, want = %v", tc.x, tc.y, got, tc.want)
		}
	}
}

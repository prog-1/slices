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
		{2, 1, [][]int{
			{1, 2},
		}},
		{4, 5, [][]int{
			{1, 2, 3, 4},
			{14, 15, 16, 5},
			{13, 20, 17, 6},
			{12, 19, 18, 7},
			{11, 10, 9, 8},
		}},
		{2, 2, [][]int{
			{1, 2},
			{4, 3},
		}},
	} {
		got := Spiral(tc.x, tc.y)
		if !equal(got, tc.want) {
			t.Errorf("Spiral(%v, %v) = %v, want = %v", tc.x, tc.y, got, tc.want)
		}
	}
}

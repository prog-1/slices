package main

import (
	"reflect"
	"testing"
)

func TestSpiral(t *testing.T) {
	for _, tc := range []struct {
		x, y int
		want [][]int
	}{
		{0, 0, [][]int{}},
		{1, 1, [][]int{
			{1},
		}},
		{3, 5, [][]int{
			{1, 2, 3},
			{12, 13, 4},
			{11, 14, 5},
			{10, 15, 6},
			{9, 8, 7},
		}},
		{5, 5, [][]int{
			{1, 2, 3, 4, 5},
			{16, 17, 18, 19, 6},
			{15, 24, 25, 21, 7},
			{14, 23, 23, 22, 8},
			{13, 12, 11, 10, 9},
		}},
	} {
		got := Spiral(tc.x, tc.y)
		if !equal(got, tc.want) {
			t.Errorf("Spiral(%v, %v) = %v, want = %v", tc.x, tc.y, got, tc.want)
		}
	}
}
func equal(a, b [][]int) bool {
	return reflect.DeepEqual(a, b)
}

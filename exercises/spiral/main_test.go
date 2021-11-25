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
		{8, 8, [][]int{
			{1, 2, 3, 4, 5, 6, 7, 8},
			{28, 29, 30, 31, 32, 33, 34, 9},
			{27, 48, 49, 50, 51, 52, 35, 10},
			{26, 47, 60, 61, 62, 53, 36, 11},
			{25, 46, 59, 64, 63, 54, 37, 12},
			{24, 45, 58, 57, 56, 55, 38, 13},
			{23, 44, 43, 42, 41, 40, 39, 14},
			{22, 21, 20, 19, 18, 17, 16, 15},
		}},
		{3, 7, [][]int{
			{1, 2, 3},
			{16, 17, 4},
			{15, 24, 5},
			{14, 23, 6},
			{13, 22, 7},
			{12, 21, 8},
			{11, 10, 9},
		}},
	} {

		if got := spiral(tc.x, tc.y); !equal(got, tc.want) {
			t.Errorf("Spiral(%v, %v) = %v, want = %v", tc.x, tc.y, got, tc.want)
		}
	}
}

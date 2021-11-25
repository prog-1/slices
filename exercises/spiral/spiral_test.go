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
		{1, 1, [][]int{
			{0},
		}},
		{2, 2, [][]int{
			{1, 2},
			{4, 3},
		}},
		{4, 9, [][]int{
			{1, 2, 3, 4},
			{22, 23, 24, 5},
			{21, 36, 25, 6},
			{20, 35, 26, 7},
			{19, 34, 27, 8},
			{18, 33, 28, 9},
			{17, 32, 29, 10},
			{16, 31, 30, 11},
			{15, 14, 13, 12},
		}},
		{8, 3, [][]int{
			{1, 2, 3, 4, 5, 6, 7, 8},
			{18, 19, 20, 21, 22, 23, 24, 9},
			{17, 16, 15, 14, 13, 12, 11, 10},
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

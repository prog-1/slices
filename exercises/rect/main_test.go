package main

import (
	"reflect"
	"testing"
)

func equal(a, b [][]int) bool {
	return reflect.DeepEqual(a, b)
}

func TestRect(t *testing.T) {
	for _, tc := range []struct {
		x, y int
		want [][]int
	}{
		{0, 0, [][]int{}},
		{1, 1, [][]int{{1}}},

		{2, 2, [][]int{{1, 1},
			{1, 1}}},

		{2, 3, [][]int{{1, 1},
			{1, 1},
			{1, 1}}},

		{3, 3, [][]int{{1, 1, 1},
			{1, 0, 1},
			{1, 1, 1}}},
		{4, 3, [][]int{{1, 1, 1, 1},
			{1, 0, 0, 1},
			{1, 1, 1, 1}}},
	} {
		got := Rect(tc.x, tc.y)
		if !equal(got, tc.want) {
			t.Errorf("Sort(%v, %v) = %v, want = %v", tc.x, tc.y, got, tc.want)
		}

	}

}

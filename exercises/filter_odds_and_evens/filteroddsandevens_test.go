package main 

import "testing"

func TestFilterEvens(t * testing.T) {
	for _, tc := range []struct {
		want []int
		s []int
	}{
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{2}, []int{2}},
		{[]int{3,6,9}, []int{3,6,9}},
		{[]int{6,4,2}, []int{6,4,2}},
	}{
		got := FilterEvens(tc.s)
		if !equal(got, tc.want) {
			t.Errorf("FilterEvens(%v) = %v, want = %v", tc.s, got, tc.want)
		}
	}
}

func equal (t, y []int) bool 

func TestFiterOdds(t * testing.T){
	
}

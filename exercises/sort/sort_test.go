package main

import (
	"math/rand"
	"testing"
	"time"
)

func Equal(a, b []int) bool {
	for k := range a {
		if a[k] != b[k] {
			return false
		}
	}
	return true
}

func TestSort(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	for i := 100; i > 0; i-- {
		tescase := make([]int, 15)
		for a := 15; a > 0; a-- {
			min := 1
			max := 100
			tescase[a-1] = rand.Intn(max-min+1) + min
		}
		got := Sort1(tescase)
		want := Sort2(tescase)
		if !Equal(got, want) {
			t.Errorf("Sort1(%v) = (%v), want = (%v)",
				tescase, got, want)
		}
	}
}

package main

import (
	"fmt"
	"sort"
)

func Sort1(s []int) []int {
	swapped := true
	for swapped {
		swapped = false
		for i := 1; i < len(s); i++ {
			if s[i-1] > s[i] {
				s[i], s[i-1] = s[i-1], s[i]
				swapped = true
			}
		}
	}
	return s
}

func Sort2(s []int) []int {
	sort.Ints(s)
	return s
}

func main() {
	fmt.Println("Enter slice size:")
	var size uint
	fmt.Scan(&size)
	input := make([]int, size)
	for i := range input {
		fmt.Scan(&input[i])
	}
	sort1 := Sort1(input)
	sort2 := Sort2(input)
	fmt.Println("Sorted using bubble-sort:", sort1)
	fmt.Println("Sorted using sort.Ints:", sort2)
}

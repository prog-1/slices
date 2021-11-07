package main

import (
	"fmt"
	"sort"
)

func minmax(a, b int) (mn, mx int) {
	if a <= b {
		return a, b
	}
	return b, a
}

func Sort1(result []int) []int {
	for range result {
		for key := range result {
			if key != 0 {
				result[key-1], result[key] = minmax(result[key-1], result[key])
			}
		}
	}
	return result
}

func Sort2(result []int) []int {
	sort.Ints(result)
	return result
}

func main() {
	fmt.Println(Sort1([]int{8, 4, 7, 5, 766, 34, 5}))
	fmt.Println(Sort2([]int{8, 4, 7, 5, 766, 34, 5}))
}

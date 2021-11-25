package main

import (
	"fmt"
	"sort"
)

func Sort1(s1 []int) {
	for i := len(s1) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if s1[j+1] < s1[j] {
				s1[j], s1[j+1] = s1[j+1], s1[j]

			}
		}
	}
}

func Sort2(s2 []int) []int {
	sort.Ints(s2)
	return s2
}

func main() {
	var count int
	fmt.Print("How many numbers do you want to enter? ")
	fmt.Scan(&count)
	var s1, s2 []int
	for i := 0; i < count; i++ {
		var n int
		fmt.Printf("Enter number #%v: ", i+1)
		fmt.Scan(&n)
		s1 = append(s1, n)
		s2 = append(s2, n)
	}
	Sort1(s1)
	Sort2(s2)
	fmt.Println("Sorted using bubble-sort:", s1)
	fmt.Println("Sorted using built-in `sort.Ints([]int)`:", s2)
}

package main

import (
	"fmt"
	"sort"
)

func Sort1(s []int) {
	len := len(s)
	for i := 0; i < len-1; i++ {
		for j := 0; j < len-i-1; j++ {
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
	}
}

func Sort2(s []int) {
	sort.Ints(s)
	return
}

func main() {
	var count int
	fmt.Print("How many numbers do you want to enter? ")
	fmt.Scan(&count)
	var s []int
	for i := 0; i < count; i++ {
		var n int
		fmt.Printf("Enter number #%v: ", i+1)
		fmt.Scan(&n)
		s = append(s, n)
	}
	Sort1(s)
	fmt.Println("Sorted (sort1):", s)
	Sort2(s)
	fmt.Println("Sorted (sort2):", s)
}

package main

import (
	"fmt"
	"sort"
)

func Sort1(s []int) []int {
	for i := len(s); i > 0; i-- {
		for j := 1; j < i; j++ {
			if s[j-1] > s[j] {
				s[j], s[j-1] = s[j-1], s[j]
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

	fmt.Println("Sorted by bublle-sort:", s)

	Sort2(s)

	fmt.Println("Sorted by built-in alg.:", s)

}

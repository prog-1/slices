package main

import (
	"fmt"
	"sort"
)

func Sort1(s []int) {
	len := len(s)
	for a := 0; a < len-1; a++ {
		for b := 0; b < len-a-1; b++ {
			if s[b] > s[b+1] {
				s[b], s[b+1] = s[b+1], s[b]
			}
		}
	}

}

func Sort2(s []int) {
	sort.Ints(s)
}

func main() {
	fmt.Println("Slice size: ")
	var size int
	fmt.Scan(&size)
	s := make([]int, size)
	for i := range s {
		fmt.Scan(&s[i])
	}
	fmt.Println("Choose what sorting algorithm you want to: \n 1 - buble sorting algorithm \n 2 - build in algorithm \n 3 - either")
	var c int
	fmt.Scan(&c)
	if c == 1 {
		Sort1(s)
		fmt.Printf("Sorted slice: %v", s)
	} else if c == 2 {
		Sort2(s)
		fmt.Printf("\nSorted slice: %v", s)
	} else if c == 3 {
		Sort1(s)
		fmt.Printf("Sorted slice: %v", s)
		Sort2(s)
		fmt.Printf("\nSorted slice: %v", s)
	} else {
		fmt.Printf("You haven't got other options. Please choose 1, 2 or 3. ^_^")
	}

}

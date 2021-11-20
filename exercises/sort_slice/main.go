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
	var y int
	fmt.Print("Enter")
	fmt.Scan(&y)
	var s []int
	for i := 0; i < y; i++ {
		var n int
		fmt.Printf("Enter number #%v: ", i+1)
		fmt.Scan(&n)
		s = append(s, n)
	}
	fmt.Println(s)
	fmt.Println("1-> custom sort")
	fmt.Println("2-> build in sort")
	var x int
	fmt.Scan(&x)
	if x == 1 {
		Sort1(s)
		fmt.Println(s)
	}
	if x == 2 {
		Sort2(s)
		fmt.Println(s)
	} else {
		fmt.Println("wrong")
	}

}

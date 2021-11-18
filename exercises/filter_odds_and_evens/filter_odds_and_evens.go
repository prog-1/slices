package main

import "fmt"

func FilterOdds(s []int) []int {
	r := make([]int, 0)
	for _, v := range s {
		if v%2 != 0 {
			r = append(r, v)
		}
	}
	return r
}

func FilterEvens(s []int) []int {
	rr := make([]int, 0)
	for _, v := range s {
		if v%2 == 0 {
			rr = append(rr, v)
		}
	}
	return rr
}

func main() {
	fmt.Println("Enter slice size:")
	var size uint
	fmt.Scan(&size)
	input := make([]int, size)
	for i := range input {
		fmt.Scan(&input[i])
	}
	fmt.Println(FilterOdds(input), "odds numbers")
	fmt.Println(FilterEvens(input), "even numbers")
}

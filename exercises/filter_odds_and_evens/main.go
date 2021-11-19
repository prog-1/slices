package main

import "fmt"

func FilterOdds(s []int) []int {
	a := make([]int, 0)
	for _, v := range s {
		if v%2 != 0 {
			a = append(a, v)
		}
	}
	return a
}

func FilterEvens(s []int) []int {
	b := make([]int, 0)
	for _, v := range s {
		if v%2 == 0 {
			b = append(b, v)
		}
	}
	return b
}

func main() {
	fmt.Println("Enter slice size:")
	var size uint
	fmt.Scan(&size)
	input := make([]int, size)
	for i := range input {
		fmt.Scan(&input[i])
	}
	fmt.Println(FilterOdds(input), "- odds numbers")
	fmt.Println(FilterEvens(input), "- evens numbers")
}

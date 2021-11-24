package main

import "fmt"

func FilterOdds(s []int) []int {
	var o []int
	for _, v := range s {
		if v%2 == 1 || v%2 == -1 {
			o = append(o, v)
		}
	}
	return o
}

func FilterEvens(s []int) []int {
	var e []int
	for _, v := range s {
		if v%2 == 0 {
			e = append(e, v)
		}
	}
	return e
}

func main() {
	fmt.Println("Enter slice size:")
	var size uint
	fmt.Scan(&size)
	input := make([]int, size)
	for i := range input {
		fmt.Scan(&input[i])
	}
	o := FilterOdds(input)
	e := FilterEvens(input)
	fmt.Println("Only odds:", o)
	fmt.Println("Only evens:", e)
}

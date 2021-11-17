package main

import "fmt"

func OnlyOdd(s []int) []int {
	a := make([]int, 0)
	for _, value := range s {
		if value%2 == 0 {
			a = append(a, value)
		}

	}
	return a
}
func OnlyEven(s []int) []int {
	b := make([]int, 0)
	for _, value := range s {
		if value%2 != 0 {
			b = append(b, value)
		}
	}
	return b
}

func main() {
	fmt.Println("Enter slice size:")
	var size int
	fmt.Scan(&size)
	input := make([]int, size)
	for i := range input {
		fmt.Scan(&input[i])
	}
	fmt.Println(OnlyEven(input))
	fmt.Println(OnlyOdd(input))
}

/*
1.Sort number to 2 groups (Odds and Evens). If number %2 == 0 even, else odd.
*/

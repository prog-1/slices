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
	a := make([]int, 0)
	for _, v := range s {
		if v%2 == 0 {
			a = append(a, v)
		}
	}
	return a
}

func main() {
	fmt.Println("Enter slice size:")
	var size uint
	fmt.Scan(&size)
	input := make([]int, size)
	for i := range input {
		fmt.Scan(&input[i])
	}
	fmt.Println("Odd numbers:", FilterOdds(input))
	fmt.Println("Even numbers:", FilterEvens(input))
}

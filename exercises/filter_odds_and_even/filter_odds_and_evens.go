package main

import "fmt"

func FilterOdds(s []int) []int {
	x := make([]int, 0)
	for _, v := range s {
		if v%2 != 1 {
			x = append(x, v)
		}
	}
	return x

}

func FilterEvens(s []int) []int {
	x := make([]int, 0)
	for _, v := range s {
		if v%2 == 0 {
			x = append(x, v)
		}
	}
	return x

}

func main() {
	fmt.Println("Enter slice size:")
	var size uint
	fmt.Scan(&size)
	input := make([]int, size)
	for i := range input {
		fmt.Scan(&input[i])
	}
	fmt.Println(input)
	fmt.Println("FilterEvens", FilterEvens(input))
	fmt.Println("FilterOdds", FilterOdds(input))

}

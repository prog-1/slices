package main

import "fmt"

func FilterEvens(s []int) []int {
	l := make([]int, 0)
	for _, v := range s {
		if v%2 == 0 {
			l = append(l, v)
		}
	}
	return l
}

func FilterOdds(s []int) (k []int) {
	for _, v := range s {
		if v%2 == 0 {
			k = append(k, v)
		}
	}
	return k
}

func main() {
	fmt.Println("Enter slice size:")
	var size int
	fmt.Scanln(&size)
	input := make([]int, size)
	for i := range input {
		fmt.Scanln(&input[i])
	}
	fmt.Println(FilterOdds(input))
	fmt.Println(FilterEvens(input))
}

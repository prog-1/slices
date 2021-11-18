package main

import "fmt"

func FilterEvens(s []int) []int {
	var filtered []int
	for _, v := range s {
		if v%2 == 0 {
			filtered = append(filtered, v)
		}
	}
	return filtered
}

func FilterOdds(s []int) []int {
	var filtered []int
	for _, v := range s {
		if v%2 != 0 {
			filtered = append(filtered, v)
		}
	}
	return filtered
}

func main() {
	fmt.Print("Enter slice size: ")
	var size int
	fmt.Scan(&size)
	s := make([]int, size)
	for i := range s {
		fmt.Scan(&s[i])
	}
	fmt.Println("Only evens:", FilterEvens(s))
	fmt.Println("Only odds:", FilterOdds(s))
}

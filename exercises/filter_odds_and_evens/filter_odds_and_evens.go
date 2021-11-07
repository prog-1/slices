package main

import "fmt"

func FilterOdds(s []int) []int {
	result := make([]int, 0)
	for _, value := range s {
		if value%2 != 0 {
			result = append(result, value)
		}
	}
	return result
}

func main() {
	fmt.Println(FilterOdds([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}))
}

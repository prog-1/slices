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

func FilterEvens(s []int) []int {
	result := make([]int, 0)
	for _, value := range s {
		if value%2 == 0 {
			result = append(result, value)
		}
	}
	return result
}

func main() {
	fmt.Println("Enter numbers(0 for stop)")
	s := make([]int, 0)
	for {
		var tmp int
		fmt.Scan(&tmp)
		if tmp == 0 {
			break
		}
		s = append(s, tmp)
	}
	fmt.Println("Odds only:", FilterOdds(s))
	fmt.Println("Evens only:", FilterEvens(s))
}

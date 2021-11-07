package main

import "fmt"

func AreOddsOnly(s []int) bool {
	for _, value := range s {
		if value%2 == 0 {
			return false
		}
	}
	return true
}

func AreEvensOnly(s []int) bool {
	for _, value := range s {
		if value%2 != 0 {
			return false
		}
	}
	return true
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
	fmt.Println("Are odd only:", AreOddsOnly(s))
	fmt.Println("Are even only:", AreEvensOnly(s))
}

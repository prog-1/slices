package main

import "fmt"

func AnyIsOdd(s []int) bool {
	for _, value := range s {
		if value%2 != 0 {
			return true
		}
	}
	return false
}

func AnyIsEven(s []int) bool {
	for _, value := range s {
		if value%2 == 0 {
			return true
		}
	}
	return false
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
	fmt.Println("Any is odd:", AnyIsOdd(s))
	fmt.Println("Any is even:", AnyIsEven(s))
}

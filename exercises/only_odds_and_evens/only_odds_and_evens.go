package main

import "fmt"

func AreOddsOnly(s []int) bool {
	for _, v := range s {
		if v%2 == 1 {
			return false
		}
	}
	return true
}

func AreEvenOnly(s []int) bool {
	for _, v := range s {
		if v%2 == 0 {
			return false
		}
	}
	return true
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
	if AreEvenOnly(input) {
		fmt.Println("All numbers evens")
	} else {
		fmt.Println("Not all numbers evens")
	}
	if AreOddsOnly(input) {
		fmt.Println("All numbers are odds")
	} else {
		fmt.Println("Not all numbers are odds")
	}
}

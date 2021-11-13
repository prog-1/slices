package main

import (
	"fmt"
)

func AreOddsOnly(s []int) bool {
	for _, v := range s {
		if v%2 == 0 {
			return false
		}
	}
	return true
}

func AreEvenOnly(s []int) bool {
	for _, v := range s {
		if v%2 != 0 {
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
	if AreOddsOnly(input) {
		fmt.Println("All numbers are odd")
	} else {
		fmt.Println("Not all numbers are odd")
	}
	if AreEvenOnly(input) {
		fmt.Println("All numbers are even")
	} else {
		fmt.Println("Not all numbers are even")
	}
}

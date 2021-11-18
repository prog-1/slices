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
		if v%2 == 1 {
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
		fmt.Println("Slice contains only numbers which are even")
	} else {
		fmt.Println("There is at least one odd number in the slice")
	}

	if AreOddsOnly(input) {
		fmt.Println("Slice contains only numbers which are odd")
	} else {
		fmt.Println("There is at least one even number in the slice")
	}
}

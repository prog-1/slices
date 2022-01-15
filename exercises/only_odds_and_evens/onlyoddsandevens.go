package main

import "fmt"

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
	fmt.Println("Enter slice size")
	var size int
	fmt.Scanln(&size)
	input := make([]int, size)
	for i := range input {
		fmt.Scanln(&input[i])
	}
	fmt.Println(input)
	if AreEvenOnly(input) {
		fmt.Println("numbers are evens")
	} else {
		fmt.Println("numbers not evens")
	}
	if AreOddsOnly(input) {
		fmt.Println("numbers are odds")
	} else {
		fmt.Println("numbers not odds")
	}
}

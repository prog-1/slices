package main

import "fmt"

func AnyIsOdd(s []int) bool {
	for _, v := range s {
		if v%2 == 1 {
			return true
		}
	}
	return false
}

func AnyIsEven(s []int) bool {
	for _, v := range s {
		if v%2 == 0 {
			return true
		}
	}
	return false
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
	if AnyIsEven(input) {
		fmt.Println("At least a single number is even")
	} else {
		fmt.Println("No even number in slice")
	}
	if AnyIsOdd(input) {
		fmt.Println("At least a single number is odd")
	} else {
		fmt.Println(" no odd number in slice")
	}

}

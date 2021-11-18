package main

import "fmt"

func AreOddsOnly(s []int) bool {
	for _, i := range s {
		if i%2 == 0 {
			return false
		}
	}
	return true
}
func AreEvenOnly(s []int) bool {
	for _, i := range s {
		if i%2 != 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("Enter slice size:")
	var size uint
	fmt.Scan(&size)
	s := make([]int, size)
	for i := range s {
		fmt.Scan(&s[i])
	}
	if AreOddsOnly(s) {
		fmt.Println("All numbers are odd")
	} else if AreEvenOnly(s) {
		fmt.Println("All numbers are even")

	} else {
		fmt.Println("There are and odd and even numbers")
	}

}

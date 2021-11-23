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
		if v%2 != 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Print("Enter slice size: ")
	var size int
	fmt.Scan(&size)
	s := make([]int, size)
	for i := range s {
		fmt.Scan(&s[i])
	}
	fmt.Println(s)
	if AreEvenOnly(s) {
		fmt.Println("All number is even")
	} else {
		fmt.Println("At least a single number in the slice is odd")
	}

	if AreOddsOnly(s) {
		fmt.Println("All numbers is odd")
	} else {
		fmt.Println("At least a single number in the slice is even")
	}

}

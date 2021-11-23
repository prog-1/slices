package main

import "fmt"

func Sort1(s []int) {
	for i := len(s) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if s[j+1] < s[j] {
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
	}
}

func main() {
	fmt.Print("Enter slice size: ")
	var size int
	fmt.Scan(&size)
	s := make([]int, size)
	for i := range s {
		fmt.Scan(&s[i])
	}
	Sort1(s)
	fmt.Println("Sorted:", s)
}

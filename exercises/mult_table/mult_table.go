package main

import "fmt"

func MultTable() (s [10][10]int) {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			s[i][j] = (i + 1) * (j + 1)
		}
	}
	return s
}

func main() {
	s := MultTable()
	for i := range s {
		for j := range s[i] {
			fmt.Printf("%4d", s[i][j])
		}
		fmt.Println()
	}
}

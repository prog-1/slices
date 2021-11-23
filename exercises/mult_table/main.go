package main

import "fmt"

func MultTable() (s [10][10]int) {
	for i := range s {
		for j := range s {
			s[i][j] = (i + 1) * (j + 1)
		}
	}
	return s
}

func main() {
	for row := range MultTable() {
		for col := range MultTable() {
			fmt.Printf("%4d", MultTable()[row][col])
		}
		fmt.Println()
	}
}

package main

import "fmt"

func Rect(x, y int) [][]int {
	s := make([][]int, y)
	for i := 0; i < y; i++ {
		s[i] = make([]int, x)
	}
	for i, j := 0, 0; j < x; j++ {
		s[i][j] = 1
	}
	for i, j := 0, 0; i < y; i++ {
		s[i][j] = 1
	}
	for i, j := y, 0; j < x; j++ {
		s[i-1][j] = 1
	}
	for i, j := 0, x; i < y; i++ {
		s[i][j-1] = 1
	}
	return s
}

func main() {
	fmt.Print("Enter width and height of a rectangle: ")
	var x, y int
	fmt.Scan(&x, &y)
	for row := 0; row < y; row++ {
		for col := 0; col < x; col++ {
			fmt.Printf("%1d", Rect(x, y)[row][col])
		}
		fmt.Println()
	}
}

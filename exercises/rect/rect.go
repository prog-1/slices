package main

import "fmt"

func Rect(x, y int) [][]int {
	s := make([][]int, x)
	for i := 0; i < x; i++ {
		s[i] = make([]int, y)
		for j := 0; j < y; j++ {
			s[0][j] = 1
			s[i][0] = 1
			s[i][y-1] = 1
			if i == x-1 {
				s[i][j] = 1
			}
		}
	}
	return s
}

func main() {
	fmt.Println("Enter height and width:")
	var h, w int
	fmt.Scan(&h, &w)
	s := Rect(h, w)
	for row := 0; row < h; row++ {
		for col := 0; col < w; col++ {
			fmt.Printf("%3d", s[row][col])
		}
		fmt.Println()
	}
}

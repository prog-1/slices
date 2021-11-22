package main

import "fmt"

func Rect(x, y int) (a [][]int) {
	s := make([][]int, x)
	for i := 0; i < x; i++ {
		s[i] = make([]int, y)
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
	for a := 0; a < y; a++ {
		for col := 0; col < x; col++ {
			fmt.Printf("%1d", Rect(x, y)[a][col])
		}
		fmt.Println()
	}
}

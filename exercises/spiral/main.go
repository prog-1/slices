package main

import "fmt"

func Spiral(x, y int) [][]int {
	s := make([][]int, y)
	for i := 0; i < y; i++ {
		s[i] = make([]int, x)
	}
	var a, b int
	for i, j := 0, 0; j < x; j++ {
		s[i][j] = 1
		a++
		if a == x {
			for ; i < y; i++ {
				s[i][j] = 1
				b++
			}
		}

	}
	return s
}

func main() {
	fmt.Print("Enter width and height of a rectangle: ")
	var x, y int
	fmt.Scan(&x, &y)
	for row := 0; row < y; row++ {
		for col := 0; col < x; col++ {
			fmt.Printf("%4d", Spiral(x, y)[row][col])
		}
		fmt.Println()
	}
}

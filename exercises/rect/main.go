package main

import "fmt"

func Rect(x, y int) [][]int {
	s := make([][]int, y)
	for row := 0; row < y; row++ {
		s[row] = make([]int, x)
	}

	for i := 0; i < x; i++ {
		s[0][i] = 1
		s[y-1][i] = 1
	}
	for i := 0; i < y; i++ {
		s[i][0] = 1
		s[i][x-1] = 1
	}
	return s
}
func main() {
	fmt.Print("Enter width and height of a rectangle: ")
	var x, y int
	fmt.Scan(&x, &y)
	if x <= 0 || y <= 0 {
		fmt.Println("Rectangle does not exist.")
		return
	}
	t := Rect(x, y)
	for i := range t {
		for j := range t[i] {
			fmt.Printf("%2d", t[i][j])
		}
		fmt.Println()
	}
}

package main

import "fmt"

func Rect(x, y int) (rect [][]int) {
	rows, cols := y, x
	r := make([][]int, rows)
	for row := 0; row < rows; row++ {
		r[row] = make([]int, cols)
	}
	for v := 0; v < x; v++ {
		r[0][v] = 1
		r[y-1][v] = 1
	}
	for v := 0; v < y; v++ {
		r[v][0] = 1
		r[v][x-1] = 1
	}
	rect = r
	return rect
}

func main() {
	fmt.Println("Enter the lenght and height of rectangle")
	var x, y int
	fmt.Scan(&x, &y)
	rectangle := Rect(x, y)
	for v := range rectangle {
		for i := range rectangle[v] {
			fmt.Println(rectangle[v][i])
			fmt.Printf("%3d", rectangle[v][i])
		}
		fmt.Println()
	}
}

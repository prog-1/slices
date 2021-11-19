package main

import "fmt"

func Rect(x, y int) [][]int {

	rect := make([][]int, y) //[y][x]int

	rect[0] = make([]int, x)
	for i := range rect[0] {
		rect[0][i] = 1
	}
	for i := 1; i < y-1; i++ {
		rect[i] = make([]int, x)
		rect[i][0] = 1
		rect[i][x-1] = 1
	}
	rect[y-1] = make([]int, x)
	for i := range rect[y-1] {
		rect[y-1][i] = 1
	}

	return rect
}

func main() {
	fmt.Println("Enter x and y (width and height)")
	var x, y int
	fmt.Scan(&x, &y)
	if x <= 0 || y <= 0 {
		fmt.Println("Width and height can't be <= 0")
	}
	r := Rect(x, y)
	for _, v := range r {
		for _, j := range v {
			fmt.Printf("%3d", j)
		}
		fmt.Println()
	}
}

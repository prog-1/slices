package main

import "fmt"

func Rect(x, y int) [][]int {
	if x <= 0 && y > 0 || y <= 0 && x > 0 {
		fmt.Println("Invalid x or y value. x or y must be > 0. Try Again ")
	}
	r := make([][]int, y)
	for i := 0; i != y && x != 0 && y != 0; i++ {
		r[i] = make([]int, x)
	}
	for a, b := 0, 0; a != y; a++ {
		r[a][b] = 1
	}
	for a, b := 0, 0; b != x; b++ {
		r[a][b] = 1
	}
	for a := 0; a != x; a++ {
		r[y-1][a] = 1
	}
	for b := 0; b != y; b++ {
		r[b][x-1] = 1
	}
	return r

}

func main() {
	var x, y int
	fmt.Println("Enter at first width, then heigt: ")
	fmt.Scan(&x, &y)
	r := Rect(x, y)
	for i := range r {
		for j := range r[i] {
			fmt.Printf("%3d", r[i][j])
		}
		fmt.Println()
	}
}

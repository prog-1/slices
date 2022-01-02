package main

import "fmt"

func Rect(x, y int) [][]int {
	a := make([][]int, y)
	for i := range a {
		a[i] = make([]int, x)
		if i == 0 || i == y-1 {
			for j := range a[i] {
				a[i][j] = 1
			}
		} else {
			a[i][0], a[i][x-1] = 1, 1
		}
	}
	return a

}
func output(a [][]int) {
	for _, row := range a {
		for _, v := range row {
			fmt.Printf("%3d", v)
		}
		fmt.Println()
	}
}

func main() {
	fmt.Println("This program prints a rectangle which borders are covered with 1.")
	fmt.Println("Write width and height of rectangle: ")
	var x, y int
	fmt.Scan(&x, &y)
	output(Rect(x, y))

}

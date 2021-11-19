package main

import "fmt"

func Spiral(x, y int) [][]int {

	matrix := make([][]int, y)

	for i := range matrix {
		matrix[i] = make([]int, x)
	}

	l := 1

	curX, curY := 0, 0
	minX, minY := 0, 0
	maxX, maxY := x-1, y-1

	for l < x*y {

		for curX <= maxX && curY <= maxY {
			matrix[curY][curX] = l
			l++
			curX++
		}
		curX--
		minY++
		curY++

		for curX <= maxX && curY <= maxY {
			matrix[curY][curX] = l
			l++
			curY++
		}
		curY--
		maxX--
		curX--

		for curX >= minX && curY >= minY {
			matrix[curY][curX] = l
			l++
			curX--
		}
		curX++
		maxY--
		curY--

		for curX >= minX && curY >= minY {
			matrix[curY][curX] = l
			l++
			curY--
		}
		curY++
		minX++
		curX++

	}
	if x == 1 && y == 1 {
		matrix[0] = []int{1}
	}

	return matrix
}

func main() {
	fmt.Println("Enter the width and hight of spiral")
	var x, y int
	fmt.Scan(&x, &y)
	if x < 0 || y < 0 {
		fmt.Println("Error. Enter only positive values")
		return
	}
	matrix := Spiral(x, y)
	for _, v := range matrix {
		for _, j := range v {
			fmt.Printf("%3d", j)
		}
		fmt.Println()
	}
}

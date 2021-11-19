package main

import "fmt"

func Spiral(x, y int) [][]int {

	spiral := make([][]int, y)

	for i := range spiral {
		spiral[i] = make([]int, x)
	}

	X, Y := 0, 0
	lf, b, rg, t := 0, 0, x-1, y-1 // lf - left, b - bottom, rg - right, t - top
	i := 1

	for i < x*y {

		for X <= rg && Y <= t {
			spiral[Y][X] = i
			i++
			X++
		}
		X--
		b++
		Y++

		for X <= rg && Y <= t {
			spiral[Y][X] = i
			i++
			Y++
		}
		Y--
		rg--
		X--

		for X >= lf && Y >= b {
			spiral[Y][X] = i
			i++
			X--
		}
		X++
		t--
		Y--

		for X >= lf && Y >= b {
			spiral[Y][X] = i
			i++
			Y--
		}
		Y++
		lf++
		X++

	}
	if x == 1 && y == 1 {
		spiral[0] = []int{1}
	}

	return spiral
}

func main() {
	fmt.Println("Enter x and y (width and height)")
	var x, y int
	fmt.Scan(&x, &y)
	spiral := Spiral(x, y)
	for _, v := range spiral {
		for _, j := range v {
			fmt.Printf("%3d", j)
		}
		fmt.Println()
	}
}

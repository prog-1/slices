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
	output(Rect(1, 2))

}

//aboba

package main

import "fmt"

func Rect(x, y int) (a [][]int) {
	rows, cols := x, y
	s := make([][]int, rows)
	for row := 0; row < rows; row++ {
		s[row] = make([]int, cols)
	}

	for i := 0; i < y; i++ {
		s[0][i] = 1
		s[x-1][i] = 1
	}
	for i := 0; i < x; i++ {
		s[i][0] = 1
		s[i][y-1] = 1
	}
	a = s
	return a
}
func main() {
	var x, y int
	fmt.Println("Enter rectangles lenght and height")
	fmt.Scan(&x, &y)
	if x < 0 || y < 0 {
		fmt.Println("Rectangles lenght and height can't be negative")
	} else {
		table := Rect(x, y)
		for i := range table {
			for j := range table[i] {
				fmt.Printf("%3d", table[i][j])
			}
			fmt.Println()
		}
	}

}

package main

import "fmt"

func Spiral(x, y int) [][]int {
	cols, rows := x, y
	s := make([][]int, rows)
	for a := 0; a < rows; a++ {
		s[a] = make([]int, cols)
	}
	i := 1
	z, g := 0, 0
	row, col, outrow, outcol := 0, 0, x-1, y-1
	for i < x*y {
		for outrow >= z && outcol >= g {
			s[g][z] = i
			i++
			z++
		}
		z--
		col++
		g++
		for outrow >= z && outcol >= g {
			s[g][z] = i
			i++
			g++
		}
		g--
		outrow--
		z--
		for row <= z && col <= g {
			s[g][z] = i
			i++
			z--
		}
		z++
		outcol--
		g--
		for row <= z && col <= g {
			s[g][z] = i
			i++
			g--
		}
		g++
		row++
		z++
	}
	return s
}

func main() {
	var x, y int
	fmt.Println("Enter width and height:")
	fmt.Scan(&x, &y)
	if x <= 0 || y <= 0 {
		fmt.Println("Rectangle does not exist")
		return
	}
	table := Spiral(x, y)
	for i := range table {
		for j := range table[i] {
			fmt.Printf("%3d", table[i][j])
		}
		fmt.Println()
	}
}

package main

import "fmt"

func spiral(x, y int) [][]int {

	s := make([][]int, y)
	for i := range s {
		s[i] = make([]int, x)
	}

	l := 1
	var a, b, xmin, ymin int

	xmax := x - 1
	ymax := y - 1
	for l < x*y {

		for a <= xmax && b <= ymax {
			s[b][a] = l
			l++

			a++
		}
		a--
		ymin++
		b++

		for a <= xmax && b <= ymax {
			s[b][a] = l
			l++

			b++
		}
		b--
		xmax--
		a--

		for a >= xmin && b >= ymin {
			s[b][a] = l
			l++

			a--
		}
		a++
		ymax--
		b--

		for a >= xmin && b >= ymin {
			s[b][a] = l
			l++

			b--
		}
		b++
		xmin++
		a++

	}
	if x == 1 && y == 1 {
		s[0] = []int{1}
	}

	return s
}

func main() {

	var x, y int
	fmt.Println("Enter rectangles lenght and height")
	fmt.Scan(&x, &y)
	if x < 0 || y < 0 {
		fmt.Println("Rectangles lenght and height can't be negative")
	} else {
		s := spiral(x, y)
		for _, v := range s {
			for _, j := range v {
				fmt.Printf("%3d", j)
			}
			fmt.Println()
		}
	}

}

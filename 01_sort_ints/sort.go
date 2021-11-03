package main

import "fmt"

func sort2(a, b int) (int, int) {
	if a < b {
		return a, b
	}
	return b, a
}

func sort4(a, b, c, d int) (int, int, int, int) {
	a, b = sort2(a, b)
	b, c = sort2(b, c)
	c, d = sort2(c, d)
	a, b = sort2(a, b)
	b, c = sort2(b, c)
	a, b = sort2(a, b)
	return a, b, c, d
}

func main() {
	fmt.Print("Enter 4 numbers: ")
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)
	a, b, c, d = sort4(a, b, c, d)
	fmt.Println("Sorted:", a, b, c, d)
}

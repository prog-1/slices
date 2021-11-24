package main

import "fmt"

func FilterOdds(s []int) []int {
	var odds []int
	for _, i := range s {
		if i%2 != 0 {
			odds = append(odds, i)
		}
	}
	fmt.Printf("Filtred Odds: %v", odds)
	return odds
}

func FilterEvens(s []int) []int {
	var evens []int
	for _, i := range s {
		if i%2 == 0 {
			evens = append(evens, i)
		}
	}
	fmt.Printf("\nFiltred Evens: %v", evens)
	return evens
}

func main() {
	fmt.Println("Enter slice size:")
	var size uint
	fmt.Scan(&size)
	input := make([]int, size)
	for i := range input {
		fmt.Scan(&input[i])
	}
	FilterOdds(input)
	FilterEvens(input)

}

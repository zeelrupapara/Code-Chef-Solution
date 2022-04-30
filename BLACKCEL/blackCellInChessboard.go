// Black cells in a chessboard
// Problem Code: BLACKCEL
// Given n (n is even), determine the number of black cells in an n×n chessboard.

// problem is solved in Golang

package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)
	numOdBlackCells := (n * n) / 2
	fmt.Println(numOdBlackCells)
}

// Blackjack
// Problem Code: BLACKJACK
// Chef is playing a variant of Blackjack, where 3 numbers are drawn and each number lies between 1 and 10 (with both 1 and 10 inclusive). Chef wins the game when the sum of these 3 numbers is exactly 21.
// Given the first two numbers A and B, that have been drawn by Chef, what should be 3-rd number that should be drawn by the Chef in order to win the game?
// Note that it is possible that Chef cannot win the game, no matter what is the 3-rd number. In such cases, report −1 as the answer

//problem sovled with Golang

package main

import "fmt"

func main() {
	var testCases int
	var results []int
	fmt.Scanf("%d", &testCases)
	for i := 1; i <= testCases; i++ {
		var x, y int
		fmt.Scanf("%d %d", &x, &y)
		result := 21 - x - y
		if result <= 10 {
			results = append(results, result)
		} else {
			results = append(results, -1)
		}
	}
	for _, o := range results {
		fmt.Println(o)
	}
}

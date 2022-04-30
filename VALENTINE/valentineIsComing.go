// Valentine is Coming
// Problem Code: VALENTINE:
// Valentine's Day is approaching and thus Chef wants to buy some chocolates for someone special.
// Chef has a total of X rupees and one chocolate costs Y rupees. What is the maximum number of chocolates Chef can buy?

//problem solved in go

package main

import "fmt"

func main() {
	var testCases int
	var results []int
	fmt.Scanf("%d", &testCases)
	for i := 1; i <= testCases; i++ {
		var x, y int
		fmt.Scanf("%d %d", &x, &y)
		results = append(results, x/y)
	}
	for _, o := range results {
		fmt.Println(o)
	}
}

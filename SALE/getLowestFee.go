// Get Lowest Free:
// Problem Code: SALE
// Chef goes to the supermarket to buy some items. Luckily there's a sale going on under which Chef gets the following offer:
// If Chef buys 3 items then he gets the item (out of those 3 items) having the lowest price as free.
// For e.g. if Chef bought 3 items with the cost 6, 2 and 4, then he would get the item with cost 2 as free. So he would only have to pay the cost of the other two items which will be 6+4=10.
// Chef buys 3 items having prices A, B and C respectively. What is the amount of money Chef needs to pay?

//problem soved with golang

package main

import "fmt"

func main() {
	var testCases int
	fmt.Scanf("%d", &testCases)
	var results []int
	for i := 1; i <= testCases; i++ {
		var a, b, c int
		fmt.Scanf("%d %d %d", &a, &b, &c)
		if a <= b && a <= c {
			results = append(results, b+c)
		} else if b <= a && b <= c {
			results = append(results, a+c)
		} else if c <= b && c <= a {
			results = append(results, a+b)
		}
	}
	for _, o := range results {
		fmt.Println(o)
	}
}

// Too many items:
// Problem Code: POLYBAGS
// Chef bought N items from a shop. Although it is hard to carry all these items in hand, so Chef has to buy some polybags to store these items.
// 1 polybag can contain at most 10 items. What is the minimum number of polybags needed by Chef?

// problem solved with Golang

package main

import "fmt"

func main() {
	var testCases int
	var results []int
	fmt.Scanf("%d", &testCases)
	for i := 1; i <= testCases; i++ {
		var numOfItems int
		fmt.Scanf("%d", &numOfItems)
		numOfBags := numOfItems / 10
		if (numOfItems % 10) != 0 {
			results = append(results, numOfBags+1)
		} else {
			results = append(results, numOfBags)
		}
	}
	for _, o := range results {
		fmt.Println(o)
	}
}

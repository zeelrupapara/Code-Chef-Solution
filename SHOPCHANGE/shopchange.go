//Shopping Change:
//Problem Code: SHOPCHANGE
// Chef went shopping and bought items worth X rupees (1≤X≤100). Unfortunately, Chef only has a single 100 rupees note.
// Since Chef is weak at maths, can you help Chef in calculating what money he should get back after paying 100 rupees for those items?

// Solution in GoLang:

package main

import "fmt"

func main() {
	var results []int
	var testCases int
	fmt.Scanf("%d", &testCases)
	for i := 1; i <= testCases; i++ {
		var x int
		fmt.Scanf("%d", &x)
		results = append(results, 100-x)
	}
	for _, o := range results {
		fmt.Println(o)
	}
}

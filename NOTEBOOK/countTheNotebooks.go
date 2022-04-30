// Count the Notebooks:
// Problem Code: NOTEBOOK
// You know that 1 kg of pulp can be used to make 1000 pages and 1 notebook consists of 100 pages.

// Suppose a notebook factory receives N kg of pulp, how many notebooks can be made from that?

//problem sloved in Golang

package main

import "fmt"

func main() {
	var testCases int
	var results []int
	fmt.Scanf("%d", &testCases)
	for i := 1; i <= testCases; i++ {
		var kgOfPulp int
		fmt.Scanf("%d", &kgOfPulp)
		numOfNotebook := (kgOfPulp * 1000) / 100
		results = append(results, numOfNotebook)
	}
	for _, o := range results {
		fmt.Println(o)
	}
}

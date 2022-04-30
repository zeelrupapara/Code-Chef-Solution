// Lazy Chef:
// Problem Code: LAZYCHF
// Read problem statements in Bengali, Mandarin Chinese, Russian, and Vietnamese as well.
// Chef is a very lazy person. Whatever work is supposed to be finished in x units of time, he finishes it in m∗x units of time. But there is always a limit to laziness, so he delays the work by at max d units of time. Given x,m,d, find the maximum time taken by Chef to complete the work.

// problem solved by Golang

package main

import "fmt"

func main() {
	var lines int
	fmt.Scanf("%d", &lines)
	var results []int
	for i := 1; i <= lines; i++ {
		var x, m, d int
		fmt.Scanf("%d %d %d", &x, &m, &d)
		if x*m < x+d {
			results = append(results, x*m)
		} else {
			results = append(results, x+d)
		}
	}
	for _, o := range results {
		fmt.Println(o)
	}
}

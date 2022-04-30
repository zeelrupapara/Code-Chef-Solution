// Cyclic Quadrilateral
// Problem Code: CYCLICQD
// Read problem statements in Mandarin Chinese, Russian, and Vietnamese as well.
// You are given the sizes of angles of a simple quadrilateral (in degrees) A, B, C and D, in some order along its perimeter. Determine whether the quadrilateral is cyclic.
// Note: A quadrilateral is cyclic if and only if the sum of opposite angles is 180∘.

//problem solve with Golang

package main

import "fmt"

func main() {
	var testCases int
	var results []string
	fmt.Scanf("%d", &testCases)
	for i := 1; i <= testCases; i++ {
		var a, b, c, d int
		fmt.Scanf("%d %d %d %d", &a, &b, &c, &d)
		if (a+c) == 180 && (b+d) == 180 {
			results = append(results, "YES")
		} else {
			results = append(results, "NO")
		}
	}
	for _, o := range results {
		fmt.Println(o)
	}
}

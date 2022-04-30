// RCB and Playoffs:
// Problem Code: RCBPLAY
// Team RCB has earned X points in the games it has played so far in this year's IPL. To qualify for the playoffs they must earn at least a total of Y points. They currently have Z games left, in each game they earn 2 points for a win, 1 point for a draw, and no points for a loss.
// Is it possible for RCB to qualify for the playoffs this year?

// problem solved with Golang

package main

import "fmt"

func main() {
	var lines int
	fmt.Scanf("%d", &lines)
	var results []string
	for i := 1; i <= lines; i++ {
		var x, y, z int
		fmt.Scanf("%d %d %d", &x, &y, &z)
		if z*2 >= (y - x) {
			results = append(results, "YES")
		} else if x >= y {
			results = append(results, "YES")
		} else {
			results = append(results, "NO")
		}
	}
	for _, o := range results {
		fmt.Println(o)
	}
}

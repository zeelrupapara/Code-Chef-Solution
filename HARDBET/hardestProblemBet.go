// Hardest Problem Bet:
// Problem Code: HARDBET
// There are 3 problems in a contest namely A,B,C respectively. Alice bets Bob that problem C is the hardest while Bob says that problem B will be the hardest.
// You are given three integers SA,SB,SC which denotes the number of successful submissions of the problems A,B,C respectively. It is guaranteed that each problem has a different number of submissions. Determine who wins the bet.
// 1) If Alice wins the bet (i.e. problem C is the hardest), then output Alice.
// 2) If Bob wins the bet (i.e. problem B is the hardest), then output Bob.
// 3) If no one wins the bet (i.e. problem A is the hardest), then output Draw.
// Note: The hardest problem is the problem with the least number of successful submissions.

// problem solved with Golang

package main

import "fmt"

func main() {
	var lines int
	fmt.Scanf("%d", &lines)
	var results []string
	for i := 1; i <= lines; i++ {
		var a, b, c int
		fmt.Scanf("%d %d %d", &a, &b, &c)
		if a < b && a < c {
			results = append(results, "Draw")
		} else if b < a && b < c {
			results = append(results, "Bob")
		} else {
			results = append(results, "Alice")
		}
	}
	for _, o := range results {
		fmt.Println(o)
	}
}

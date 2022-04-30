// Chef and NextGen Problem
// Code: HELIUM3

// Chef is currently working for a secret research group called NEXTGEN. While the rest of the world is still in search of a way to utilize Helium-3 as a fuel, NEXTGEN scientists have been able to achieve 2 major milestones:

// 1.Finding a way to make a nuclear reactor that will be able to utilize Helium-3 as a fuel
// 2.Obtaining every bit of Helium-3 from the moon's surface
// Moving forward, the project requires some government funding for completion, which comes under one condition: to prove its worth, the project should power Chefland by generating at least A units of power each year for the next B years.

// Help Chef determine whether the group will get funded assuming that the moon has X grams of Helium-3 and 1 gram of Helium-3 can provide Y units of power.

//problem solved in go

package main

import "fmt"

func main() {
	var testCases int
	var results []string
	fmt.Scanf("%d", &testCases)
	for i := 1; i <= testCases; i++ {
		var a, b, x, y int
		fmt.Scanf("%d %d %d %d", &a, &b, &x, &y)
		if a*b <= x*y {
			results = append(results, "YES")
		} else {
			results = append(results, "NO")
		}
	}
	for _, o := range results {
		fmt.Println(o)
	}
}

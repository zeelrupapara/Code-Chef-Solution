// Car or Bus:
// Problem Code: TRAVELFAST
// Chef wants to reach home as soon as possible. He has two options:
// Travel with his BIKE which takes X minutes.
// Travel with his CAR which takes Y minutes.
// Which of the two options is faster or do they both take same time?

//problem solved in Golang

package main

import "fmt"

func main() {
	var testCases int
	var results []string
	fmt.Scanf("%d", &testCases)
	for i := 1; i <= testCases; i++ {
		var bike, car int
		fmt.Scanf("%d %d", &bike, &car)
		if bike < car {
			results = append(results, "BIKE")
		} else if car < bike {
			results = append(results, "CAR")
		} else {
			results = append(results, "SAME")
		}
	}
	for _, o := range results {
		fmt.Println(o)
	}
}

// Pass or Fail
// Problem Code: PASSORFAIL
// Chef is struggling to pass a certain college course.
// The test has a total of N question, each question carries 3 marks for a correct answer and −1 for an incorrect answer. Chef is a risk-averse person so he decided to attempt all the questions. It is known that Chef got X questions correct and the rest of them incorrect. For Chef to pass the course he must score at least P marks.
// Will Chef be able to pass the exam or not?

//problem solved with Golang

package main

import "fmt"

func main() {
	var testCases int
	var results []string
	fmt.Scanf("%d", &testCases)
	for i := 1; i <= testCases; i++ {
		var x, y, p int
		fmt.Scanf("%d %d %d", &x, &y, &p)
		result := ((x * 3) - (x - y)) - ((x - y) * 3)
		if result >= p {
			results = append(results, "PASS")
		} else {
			results = append(results, "FAIL")
		}
	}
	for _, o := range results {
		fmt.Println(o)
	}
}

// Debug
// input: 15 10 10
		// ((x * 3) - (x - y)) - ((x - y) * 3)
		// ((15*3)-(15-10))-((15-10)*3)
		// ((45)-(5))-((5)*3)
		// (40)-(15)
		// (25) > 10 // true
		// Print: PASS

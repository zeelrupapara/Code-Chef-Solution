package main

import (
	"fmt"
	"math"
)

var (
	N     int
	line  int
	value int
	abs   int
)

func main() {
	fmt.Scanf("%d", &N)
	fmt.Scanf("%d", &line)
	for i := 1; i <= line; i++ {
		fmt.Scanf("%d", &value)
		if value < N+2 || value > N*3 {
			fmt.Println(0)
		} else {
			float := float64(2*N + 1 - value)
			abs = int(math.Abs(float))
			fmt.Println(N - abs)
		}
	}
}

package main

import (
	"fmt"
	"math"
	"sort"
)

func Part1() {
	result = 0.0

	sort.Ints(left)
	sort.Ints(right)

	for i, v := range left {
		result += math.Abs(float64(v - right[i]))
	}

	fmt.Println("Result for part 1: ", int(result))
}

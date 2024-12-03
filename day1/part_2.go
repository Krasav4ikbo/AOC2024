package main

import (
	"fmt"
)

func Part2() {
	result = 0.0

	for _, v := range left {
		count := 0
		for _, vR := range right {
			if vR == v {
				count++
			}
		}
		result += float64(count * v)
	}

	fmt.Println("Result for part 2: ", int(result))
}

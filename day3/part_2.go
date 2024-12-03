package main

import (
	"strings"
)

func Part2() {
	donts := strings.Split(FullText, "don't()")
	PartTwoResult += Multiply(donts[0])
	donts = donts[1:]
	for _, dont := range donts {
		dos := strings.Split(dont, "do()")
		for i, do := range dos {
			if i > 0 {
				PartTwoResult += Multiply(do)
			}
		}
	}
}

package main

import (
	"strings"
)

func Part2(line string) {
	if IsSafe(line) {
		PartTwoResult++
	} else {
		numbers := strings.Split(line, " ")
		for i, _ := range numbers {
			var tempSlice []string
			for j, number := range numbers {
				if i != j {
					tempSlice = append(tempSlice, number)
				}
			}

			if IsSafe(strings.Join(tempSlice, " ")) {
				PartTwoResult++
				return
			}
		}
	}
}

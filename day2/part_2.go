package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewReader(file)

	var result = 0
	var numbers []string
	for line, err := scanner.ReadString('\n'); err == nil; line, err = scanner.ReadString('\n') {
		numbers = strings.Split(strings.TrimSpace(line), " ")
		unsafe := findAllUnsafe(numbers)
		if len(unsafe) == 0 {
			result++
		} else {
			var removeCurrent []string
			var removePrevious []string
			var removeNext []string
			for _, unsafeValue := range unsafe {
				for index, number := range numbers {
					if index != unsafeValue {
						removeCurrent = append(removeCurrent, number)
					}

					if index != unsafeValue-1 {
						removePrevious = append(removePrevious, number)
					}

					if index != unsafeValue+1 {
						removeNext = append(removeNext, number)
					}
				}

				unsafe = findAllUnsafe(removeCurrent)
				if len(unsafe) == 0 {
					result++
				} else {
					unsafe = findAllUnsafe(removePrevious)
					if len(unsafe) == 0 {
						result++
					} else {
						unsafe = findAllUnsafe(removeNext)
						if len(unsafe) == 0 {
							result++
						}
					}
				}
			}

		}
	}

	fmt.Println(result)
}

func findAllUnsafe(numbers []string) []int {
	var unsafe []int
	current, _ := strconv.Atoi(numbers[0])
	next, _ := strconv.Atoi(numbers[1])
	diff := current - next
	var path bool

	if diff < 0 {
		path = true
	} else {
		path = false
	}

	for index, number := range numbers {
		if index != len(numbers)-1 {
			current, _ = strconv.Atoi(number)
			next, _ = strconv.Atoi(numbers[index+1])
			diff = current - next

			if diff > 0 && path ||
				diff < 0 && !path ||
				diff == 0 ||
				math.Abs(float64(diff)) > 3 {
				unsafe = append(unsafe, index)
			}
		}
	}

	return unsafe
}

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

	for line, err := scanner.ReadString('\n'); err == nil; line, err = scanner.ReadString('\n') {
		if isSafe(strings.TrimSpace(line)) {
			result++
		}
	}

	fmt.Println(result)
}

func isSafe(row string) bool {
	numbers := strings.Split(row, " ")
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
				return false
			}
		}
	}

	return true
}

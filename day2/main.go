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

var (
	PartOneResult int
	PartTwoResult int
)

func readFileByLine() {
	file, err := os.Open("./day2/input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewReader(file)

	for line, err := scanner.ReadString('\n'); err == nil; line, err = scanner.ReadString('\n') {
		Part1(strings.TrimSpace(line))
		Part2(strings.TrimSpace(line))
	}
}

func IsSafe(line string) bool {
	numbers := strings.Split(line, " ")
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

func main() {
	readFileByLine()
	fmt.Println("Result for part 1: ", PartOneResult)
	fmt.Println("Result for part 2: ", PartTwoResult)
}

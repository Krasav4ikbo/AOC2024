package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	PartOneResult int
	PartTwoResult int
	FullText      string
)

func readFileByLine() {
	file, err := os.Open("./day3/input")
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	scanner := bufio.NewReader(file)

	for line, err := scanner.ReadString('\n'); err == nil; line, err = scanner.ReadString('\n') {
		FullText += strings.TrimSpace(line)
	}
}

func Multiply(text string) int {
	re, _ := regexp.Compile("mul\\(\\d{1,3},\\d{1,3}\\)")
	matches := re.FindAllStringSubmatch(text, -1)
	result := 0
	for _, match := range matches {
		match[0] = strings.Replace(match[0], "mul(", "", -1)
		match[0] = strings.Replace(match[0], ")", "", -1)
		numbers := strings.Split(match[0], ",")
		first, _ := strconv.Atoi(numbers[0])
		second, _ := strconv.Atoi(numbers[1])

		result += first * second
	}

	return result
}

func main() {
	readFileByLine()
	Part1()
	Part2()
	fmt.Println("Result for part 1: ", PartOneResult)
	fmt.Println("Result for part 2: ", PartTwoResult)
}

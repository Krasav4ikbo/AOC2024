package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

var result = 0.0
var left []int
var right []int

func readFileByLine() {
	file, err := os.Open("./day1/input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewReader(file)

	for line, err := scanner.ReadString('\n'); err == nil; line, err = scanner.ReadString('\n') {
		CheckLine(strings.TrimSpace(line))
	}
}

func CheckLine(line string) {
	row := strings.Split(strings.TrimSpace(line), "   ")
	numLeft, _ := strconv.Atoi(row[0])
	numRight, _ := strconv.Atoi(row[1])
	left = append(left, numLeft)
	right = append(right, numRight)
}

func main() {
	readFileByLine()
	Part2()
	Part1()
}

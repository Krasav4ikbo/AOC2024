package main

import (
	"bufio"
	"fmt"
	"log"
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
	var left []int
	var right []int

	for line, err := scanner.ReadString('\n'); err == nil; line, err = scanner.ReadString('\n') {
		row := strings.Split(strings.TrimSpace(line), "   ")
		numLeft, _ := strconv.Atoi(row[0])
		numRight, _ := strconv.Atoi(row[1])
		left = append(left, numLeft)
		right = append(right, numRight)
	}

	for _, v := range left {
		count := 0
		for _, vR := range right {
			if vR == v {
				count++
			}
		}
		result += count * v
	}

	fmt.Println(result)
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
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

	var result = 0.0
	var left []int
	var right []int

	for line, err := scanner.ReadString('\n'); err == nil; line, err = scanner.ReadString('\n') {
		row := strings.Split(strings.TrimSpace(line), "   ")
		numLeft, _ := strconv.Atoi(row[0])
		numRight, _ := strconv.Atoi(row[1])
		left = append(left, numLeft)
		right = append(right, numRight)
	}

	sort.Ints(left)
	sort.Ints(right)

	for i, v := range left {
		result += math.Abs(float64(v - right[i]))
	}

	fmt.Println(int(result))
}

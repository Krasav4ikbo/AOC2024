package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	PartOneResult int
	PartTwoResult int
	Table         map[int]map[int]string
	Letters       []string
)

func readFileByLine() {
	file, err := os.Open("./day4/input")
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
	Table = map[int]map[int]string{}
	Letters = []string{"X", "M", "A", "S"}
	i := 0
	for line, err := scanner.ReadString('\n'); err == nil; line, err = scanner.ReadString('\n') {
		if Table[i] == nil {
			Table[i] = make(map[int]string)
		}
		array := strings.Split(strings.TrimSpace(line), "")
		for j, arr := range array {
			Table[i][j] = arr
		}
		i++
	}
}

func main() {
	readFileByLine()
	Part1()
	Part2()
	fmt.Println("Result for part 1: ", PartOneResult)
	fmt.Println("Result for part 2: ", PartTwoResult)
}

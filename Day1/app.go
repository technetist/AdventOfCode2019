package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	filepath := "/Users/adrien/go/src/github.com/vehcklox/AdventOfCode2019/Day1/data-set.txt"
	total := 0.00
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		value, err := strconv.ParseFloat(scanner.Text(), 10)
		if err != nil {
			panic(err)
		}
		tempValue := (math.Floor(value / 3)) -2
		total += tempValue
		fmt.Printf("%.1f\n", total)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

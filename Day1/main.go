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
	t := 0.00
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s, err := strconv.ParseFloat(scanner.Text(), 10)
		if err != nil {
			panic(err)
		}
		f := (math.Floor(s / 3)) -2
		t += f
	}

	fmt.Printf("%.1f\n", t)

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

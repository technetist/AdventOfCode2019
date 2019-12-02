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
		value, err := strconv.ParseFloat(scanner.Text(), 10)
		if err != nil {
			panic(err)
		}
		u, d := 0.00, math.Floor(value/3)
		s := d - 2
		f := s
		r := s
		for {
			if math.Floor(f/3) <= 0.00 {
				u = r
				break
			} else {
				dividedValue := math.Floor(f / 3)
				subtractedValue := dividedValue - 2
				if subtractedValue <= 0.00 {
					u = r
					break
				} else {
					f = subtractedValue
					r += f
				}
			}
		}
		t += u
	}

	fmt.Printf("%.1f\n", t)

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

package main

import (
	"fmt"
)

func part1(data []int) int {

	counter := 0

	for i, v := range data {
		if i == 0 {
			//fmt.Println(v, " (N/A - no previous measurement)")
		} else if v > data[i-1] {
			//fmt.Println(v, " (increased)")
			counter++
		} else {
			//fmt.Println(v, " (decreased)")
		}
	}

	return counter
}

func part2(data []int) int {

	counter := 0

	for i := range data[0 : len(data)-3] {

		A := sum(data[i : i+3])
		B := sum(data[i+1 : i+4])

		//fmt.Println(A, data[i:i+3], " vs ", B, data[i+1:i+4])

		if B > A {
			counter++
		}
	}

	return counter
}

func main() {

	data, err := readIntegerLines("./input")
	if err != nil {
		fmt.Println("Cannot read file", err)
	}

	fmt.Println("Part 1 result: ", part1(data))
	fmt.Println("Part 2 result: ", part2(data))
}

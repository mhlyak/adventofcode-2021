package main

import (
	"bufio"
	"os"
	"strconv"
)

func readIntegerLines(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var lines []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		x, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		lines = append(lines, x)
	}
	return lines, scanner.Err()
}

func sum(data []int) int {
	sum := 0
	for _, v := range data {
		sum += v
	}
	return sum
}

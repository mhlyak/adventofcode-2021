package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Position struct {
	Horizontal int
	Depth      int
	Aim        int
}

func (p *Position) multiply() int {
	return p.Horizontal * p.Depth
}

type Command struct {
	Type  string
	Units int
}

func readCommandLines(path string) ([]Command, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var lines []Command
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), " ")
		units, _ := strconv.Atoi(s[1])
		lines = append(lines, Command{s[0], units})
	}
	return lines, scanner.Err()
}

package main

import "fmt"

func part1(commands []Command) Position {
	pos := Position{0, 0, 0}
	for _, v := range commands {
		switch v.Type {
		case "forward":
			pos.Horizontal += v.Units
		case "down":
			pos.Depth += v.Units
		case "up":
			pos.Depth -= v.Units
		}
	}
	return pos
}

func part2(commands []Command) Position {
	pos := Position{0, 0, 0}
	for _, v := range commands {
		switch v.Type {
		case "forward":
			pos.Horizontal += v.Units
			pos.Depth += (pos.Aim * v.Units)
		case "down":
			pos.Aim += v.Units
		case "up":
			pos.Aim -= v.Units
		}
	}
	return pos
}

func main() {

	data, _ := readCommandLines("./input")

	part1pos := part1(data)
	fmt.Println(part1pos, part1pos.multiply())

	part2pos := part2(data)
	fmt.Println(part2pos, part2pos.multiply())
}

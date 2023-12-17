package main

import (
	"fmt"
	"strconv"
	"strings"
	"utils"
)

func Part1(filename string) int {
	result := 0
	lines := utils.ReadInput(filename)

	for _, line := range lines {
		idGame := strings.Split(line, ": ")
		id, _ := strconv.Atoi(idGame[0][5:])
		game := idGame[1]
		valid := true
		for _, turn := range strings.Split(game, "; ") {
			for _, num_color := range strings.Split(turn, ", ") {
				num_color_array := strings.Split(num_color, " ")
				num, _ := strconv.Atoi(num_color_array[0])
				color := strings.TrimSpace(strings.Join(num_color_array[1:], ""))
				if num > 14 || (color == "red" && num > 12) || (color == "green" && num > 13) {
					valid = false
					break
				}
			}
			if !valid {
				break
			}
		}
		if valid {
			result += id
		}

	}

	return result
}

func Part2(filename string) int {
	result := 0
	lines := utils.ReadInput(filename)

	for _, line := range lines {
		idGame := strings.Split(line, ": ")
		game := idGame[1]

		maxr, maxb, maxg := 0, 0, 0

		for _, turn := range strings.Split(game, "; ") {
			for _, num_color := range strings.Split(turn, ", ") {
				num_color_array := strings.Split(num_color, " ")
				num, _ := strconv.Atoi(num_color_array[0])
				color := strings.TrimSpace(strings.Join(num_color_array[1:], ""))

				if color == "red" {
					maxr = utils.Max(num, maxr)
				} else if color == "blue" {
					maxb = utils.Max(num, maxb)
				} else {
					maxg = utils.Max(num, maxg)
				}
			}
		}

		result += maxr * maxg * maxb
	}

	return result
}

func main() {
	part1 := Part1("./day_2/input_1.txt")
	fmt.Printf("part 1: %v", part1)
}

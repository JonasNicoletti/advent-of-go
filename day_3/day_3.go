package main

import (
	"strconv"
	"strings"
	"utils"
)

func Part1(filename string) int {
	result := 0
	lines := utils.ReadInput(filename)
	width, height := len(lines[0]), len(lines)
	for r, row := range lines {
		c := 0
		for c < width {
			for c < width && !utils.IsDigit(string(row[c])) {
				c += 1
			}
			if c >= width {
				break
			}
			hasAdjSimbol, _ := FindAdjSimbols(lines, r, c, height, width)
			if hasAdjSimbol {
				n, _ := ExtractNumber(row, c, width)
				result += n
			}

			for c < width && utils.IsDigit(string(row[c])) {
				c += 1
			}
		}
	}
	return result
}

func Part2(filename string) int {
	result := 0
	lines := utils.ReadInput(filename)
	width, height := len(lines[0]), len(lines)
	gears := make(map[utils.Coordinate][]int)
	for r, row := range lines {
		c := 0
		for c < width {
			for c < width && !utils.IsDigit(string(row[c])) {
				c += 1
			}
			if c >= width {
				break
			}
			hasAdjSimbol, adjGears := FindAdjSimbols(lines, r, c, height, width)
			if hasAdjSimbol {
				n, _ := ExtractNumber(row, c, width)
				for _, gear := range adjGears {
					gears[gear] = append(gears[gear], n)
				}
			}

			for c < width && utils.IsDigit(string(row[c])) {
				c += 1
			}
		}
	}
	for _, gear := range gears {
		if len(gear) == 2 {
			result += gear[0] * gear[1]
		}
	}
	return result
}

func ExtractNumber(text string, start int, length int) (int, error) {
	end := start + 1
	for end < length && utils.IsDigit(string(text[end])) {
		end += 1
	}
	return strconv.Atoi(text[start:end])
}

func FindAdjSimbols(grid []string, r int, start int, height int, width int) (bool, []utils.Coordinate) {
	start = utils.Max(0, start-1)
	row := grid[r]
	var gears []utils.Coordinate
	hasAdjSimbol := false
	for i := start; i < width; i++ {
		//check above
		if r > 0 && !strings.Contains("0123456789.", string(grid[r-1][i])) {
			hasAdjSimbol = true
			if string(grid[r-1][i]) == "*" {
				gears = append(gears, utils.Coordinate{X: r - 1, Y: i})
			}
		}
		// check below
		if r < height-1 && !strings.Contains("0123456789.", string(grid[r+1][i])) {
			hasAdjSimbol = true
			if string(grid[r+1][i]) == "*" {
				gears = append(gears, utils.Coordinate{X: r + 1, Y: i})
			}
		}

		if !utils.IsDigit(string(row[i])) {
			hasAdjSimbol = hasAdjSimbol || string(row[i]) != "."
			if string(grid[r][i]) == "*" {
				gears = append(gears, utils.Coordinate{X: r, Y: i})
			}

			if i > start {
				break
			}
		}
	}
	return hasAdjSimbol, gears
}

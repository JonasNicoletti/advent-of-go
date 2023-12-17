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
		fmt.Printf("%v \n", id)
		for _, turn := range strings.Split(game, "; ") {
			fmt.Printf("%v \n", turn)
			for _, num_color := range strings.Split(turn, ", ") {
				num_color_array := strings.Split(num_color, " ")
				fmt.Printf("%v \n", num_color_array)
				num, _ := strconv.Atoi(num_color_array[0])
				color := strings.TrimSpace(strings.Join(num_color_array[1:], ""))
				if num > 14 || (color == "red" && num > 12) || (color == "green" && num > 13) {
					fmt.Printf("\n\n")
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

func main() {
	part1 := Part1("./day_2/input_1.txt")
	fmt.Printf("part 1: %v", part1)
}

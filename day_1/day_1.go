package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func ReadInput(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("error opening file with %v", err)
		return nil
	}

	fileScanner := bufio.NewScanner(file)

	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	file.Close()
	return fileLines
}

func FindFirstAndLastPart1(text string) int {
	r, _ := regexp.Compile("[0-9]")
	matchFirst := false
	matchLast := false
	firstChar := ""
	lastChar := ""
	for i := range text {
		if !matchFirst {
			firstChar = string(text[i])
			matchFirst = r.MatchString(firstChar)
		}
		if !matchLast {
			lastChar = string(text[len(text)-1-i])
			matchLast = r.MatchString(lastChar)
		}

		if matchFirst && matchLast {
			res, _ := strconv.Atoi(firstChar + lastChar)
			return res
		}
	}
	return 0
}

func MatchNumberOrString(text string) (int, string) {
	r, _ := regexp.Compile("([0-9]|one|two|three|four|five|six|seven|eight|nine)")
	indexes := r.FindStringIndex(text)
	if len(indexes) > 0 {
		return r.FindStringIndex(text)[0], r.FindString(text)
	}
	return 0, r.FindString(text)
}

func isDigit(text string) (int, error) {
	rIsNumber, _ := regexp.Compile("[0-9]")
	stringToNumber := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9}

	if rIsNumber.MatchString(text[0:1]) {
		return strconv.Atoi(text[0:1])
	}
	for str, num := range stringToNumber {
		if strings.HasPrefix(text, str) {
			return num, nil
		}
	}
	return -1, errors.New("no found")
}

func Part1(fileName string) int {
	result := 0

	lines := ReadInput(fileName)
	for _, line := range lines {
		result += FindFirstAndLastPart1(line)
	}

	return result
}

func Part2(fileName string) int {
	total := 0

	lines := ReadInput(fileName)
	firstDigit := -1
	secondDigit := -1
	for _, line := range lines {
		for i := 0; i < len(line); i++ {
			firstDigit, _ = isDigit(line[i:])
			if firstDigit > -1 {
				break
			}
		}
		for i := len(line) - 1; i > -1; i-- {
			secondDigit, _ = isDigit(line[i:])
			if secondDigit > -1 {
				break
			}
		}
		total += 10*firstDigit + secondDigit
	}

	return total
}

func main() {
	part2 := Part2("input_part2_test.txt")
	fmt.Printf("part 2: %v", part2)
}

package day_1

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
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

func MatchNumberOrString(text string) string {
	r, _ := regexp.Compile("([0-9]|one|two|three|four|five|six|seven|eight|nine)")
	return r.FindString(text)
}

func FindFirstAndLastPart2(text string) int {
	rIsNumber, _ := regexp.Compile("[0-9]")
	stringToNumber := map[string]string{"one": "1", "two": "2", "three": "3", "four": "4", "five": "5", "six": "6", "seven": "7", "eight": "8", "nine": "9"}

	matchFirst := ""
	matchLast := ""
	for i := range text {
		if matchFirst == "" {
			matchFirst = MatchNumberOrString(string(text[0:i]))
		}
		if matchLast == "" {
			matchLast = MatchNumberOrString(string(text[len(text)-1-i:]))
		}

		if matchFirst != "" && matchLast != "" {
			if !rIsNumber.MatchString(matchFirst) {
				matchFirst = stringToNumber[matchFirst]
			}
			if !rIsNumber.MatchString(matchLast) {
				matchLast = stringToNumber[matchLast]
			}
			res, _ := strconv.Atoi(matchFirst + matchLast)
			return res
		}
	}
	return 0
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
	result := 0

	lines := ReadInput(fileName)
	for _, line := range lines {
		result += FindFirstAndLastPart2(line)
	}

	return result
}

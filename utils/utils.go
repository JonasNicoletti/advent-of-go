package utils

import (
	"bufio"
	"log"
	"os"
	"regexp"
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

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func IsDigit(s string) bool {
	rexIsNumber, _ := regexp.Compile("[0-9]")

	return rexIsNumber.MatchString(s)
}

type Coordinate struct {
	X int
	Y int
}

package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	part1()
	part2()
}

func part1() {
	in, err := os.ReadFile("input.txt")
	if err != nil {
		return
	}

	lines := strings.Split(string(in), "\n")

	sum := 0

	for _, line := range lines {
		var firstDigit string
		var lastDigit string

		// search forward
		for _, i := range line {
			// rune comparison, since iterating over str gives rune
			if unicode.IsDigit(i) {
				firstDigit = string(i)
				break
			}
		}

		// search backwards
		for j := len(line) - 1; j >= 0; j-- {
			_, err := strconv.Atoi(string(line[j]))
			if err == nil {
				lastDigit = string(line[j])
				break
			}
		}

		doubleDigit, err := strconv.Atoi(fmt.Sprintf("%s%s", firstDigit, lastDigit))
		if err != nil {
			log.Fatalf("failed to convert string to int")
		}

		sum += doubleDigit
	}
	fmt.Println("Part 1 - sum:", sum)
}

var numberNames = []string{
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
}

func getNumberNamePrefix(in string) string {
	for _, numberName := range numberNames {
		if strings.HasPrefix(in, numberName) {
			return numberName
		}
	}
	return ""
}

var numberNameToDigitString = map[string]string{
	// note: input has no 0/zero cases
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func part2() {
	in, err := os.ReadFile("input.txt")
	if err != nil {
		return
	}

	lines := strings.Split(string(in), "\n")

	sum := 0

	for _, line := range lines {
		var firstDigit string
		var lastDigit string

		// search forward
		for i := 0; i <= len(line)-1; i++ {
			_, err := strconv.Atoi(string(line[i]))
			if err == nil {
				firstDigit = string(line[i])
				break
			}
			numberName := getNumberNamePrefix(line[i:])
			if numberName != "" {
				firstDigit = numberNameToDigitString[numberName]
				break
			}
		}

		// search backwards
		for j := len(line) - 1; j >= 0; j-- {
			_, err := strconv.Atoi(string(line[j]))
			if err == nil {
				lastDigit = string(line[j])
				break
			}
			numberName := getNumberNamePrefix(line[j:])
			if numberName != "" {
				lastDigit = numberNameToDigitString[numberName]
				break
			}
		}

		doubleDigit, err := strconv.Atoi(fmt.Sprintf("%s%s", firstDigit, lastDigit))
		if err != nil {
			log.Fatalf("failed to convert string to int")
		}

		sum += doubleDigit
	}
	fmt.Println("Part 2 - sum:", sum)
}

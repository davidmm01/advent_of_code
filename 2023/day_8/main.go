package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strings"
	"time"
)

var letterRegex = regexp.MustCompile(`[A-Z0-9]{3}`)

type direction map[string]string

type directionMap map[string]direction

func main() {
	directionSequence, directionMap := processInput()
	part1(directionSequence, directionMap)

	start := time.Now()
	part2(directionSequence, directionMap)
	fmt.Printf("part2 took %s\n", time.Since(start))
}

func processInput() (string, directionMap) {
	in, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal("couldn't even read the file dave wtf")
	}

	lines := strings.Split(string(in), "\n")

	directionSequence := lines[0]

	lines = lines[2:]

	directionMap := make(directionMap)

	for _, line := range lines {
		submatches := letterRegex.FindAllStringSubmatch(line, -1)
		direction := make(direction)
		direction["L"] = submatches[1][0]
		direction["R"] = submatches[2][0]

		directionMap[submatches[0][0]] = direction
	}
	return directionSequence, directionMap
}

func part1(directionSequence string, directionMap directionMap) {
	// always start at AAA
	currentElement := "AAA"
	targetElement := "ZZZ"
	stepsPerformed := 0

	for currentElement != targetElement {
		direction := string(directionSequence[stepsPerformed%len(directionSequence)])
		currentElement = directionMap[currentElement][direction]
		stepsPerformed++
	}

	fmt.Println("Part 1:", stepsPerformed)
}

func part2(directionSequence string, directionMap directionMap) {
	startingPositions := []string{}
	for key := range directionMap {
		if strings.HasSuffix(key, "A") {
			startingPositions = append(startingPositions, key)
		}
	}

	// It's all based off identifying that u will land on a position ending in Z in cycles

	// Observe
	// startingPosition: PTA
	// arrives on a location ending in Z after this many operatios: [12361 24722 37083 49444 61805 74166 86527 98888 111249 123610 135971]
	// startingPosition: CRA
	// arrives on a location ending in Z after this many operatios: [16043 32086 48129 64172 80215 96258 112301 128344 144387 160430 176473]
	// startingPosition: BGA
	// arrives on a location ending in Z after this many operatios: [18673 37346 56019 74692 93365 112038 130711 149384 168057 186730 205403]
	// startingPosition: AAA
	// arrives on a location ending in Z after this many operatios: [13939 27878 41817 55756 69695 83634 97573 111512 125451 139390 153329]
	// startingPosition: DVA
	// arrives on a location ending in Z after this many operatios: [11309 22618 33927 45236 56545 67854 79163 90472 101781 113090 124399]
	// startingPosition: JQA
	// arrives on a location ending in Z after this many operatios: [19199 38398 57597 76796 95995 115194 134393 153592 172791 191990 211189]

	// So we observe that from each starting position, they return to a location ending in Z in a cycle of operations.

	cycleLengths := []int{}
	for _, startingPosition := range startingPositions {
		currentPosition := startingPosition

		count := 0
		first := true

		for currentPosition != startingPosition || first {
			direction := string(directionSequence[count%len(directionSequence)])
			currentPosition = directionMap[currentPosition][direction]
			first = false
			count++
			if strings.HasSuffix(currentPosition, "Z") {
				cycleLengths = append(cycleLengths, count)
				break
			}
		}
	}

	// we can find the lowest number of steps that each path ends in Z if we find the lowest common denominator of the cycle numbers
	// cycle lengths: [18673 11309 19199 13939 16043 12361]

	highestPowerPrimeMap := make(primeMap)
	for _, length := range cycleLengths {
		// to do this, calculate a prime factorization map of each number, which has
		// keys == prime numbers, value == number of occurances
		primeMap := primeFactors(length)

		// then, compare the highestPower prime map against the current primeMap, and let any new highest powers
		// be set in the highestPowerPrimeMap
		for key, value := range primeMap {
			if primeMap[key] > highestPowerPrimeMap[key] {
				highestPowerPrimeMap[key] = value
			}
		}
	}

	// now do the final math
	// take each prime to its highest observed power and multiply them all together
	total := 1
	for key, value := range highestPowerPrimeMap {
		total = total * int(math.Pow(float64(key), float64(value)))
	}

	fmt.Println("Part 2:", total)
}

func allPositionsEndInZ(positions []string) bool {
	for _, position := range positions {
		if !strings.HasSuffix(position, "Z") {
			return false
		}
	}
	return true
}

type primeMap map[int]int

// inspo from https://www.geeksforgeeks.org/print-all-prime-factors-of-a-given-number/
// just added a map
func primeFactors(number int) primeMap {
	var primeMap = make(primeMap)
	for number%2 == 0 {
		primeMap[2]++
		number = number / 2
	}

	for i := 3; i <= int(math.Sqrt(float64(number))); i += 2 {
		for number%i == 0 {
			primeMap[i]++
			number = number / i
		}
	}

	if number > 2 {
		primeMap[number]++
	}

	return primeMap
}

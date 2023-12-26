package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

var letterRegex = regexp.MustCompile(`[A-Z0-9]{3}`)

type direction map[string]string

type directionMap map[string]direction

func main() {
	directionSequence, directionMap := processInput()
	part1(directionSequence, directionMap)

	// part 2 doesn't work yet, its too slow

	// start := time.Now()
	// part2(directionSequence, directionMap)
	// fmt.Printf("part2 took %s", time.Since(start))
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

	currentPositions := startingPositions
	stepsPerformed := 0

	for !allPositionsEndInZ(currentPositions) {
		// fmt.Println("starting new step")
		direction := string(directionSequence[stepsPerformed%len(directionSequence)])
		// fmt.Println(" direction:", direction)
		for i := 0; i < len(currentPositions); i++ {
			// fmt.Println("  i:", i)
			currentPositions[i] = directionMap[currentPositions[i]][direction]

		}
		stepsPerformed++
		// fmt.Println("currentPositions", currentPositions)
		// if stepsPerformed > 10000 {
		// 	return
		// }
	}

	fmt.Println("Part 2:", stepsPerformed)
}

// is allPositionsEndInZ adding too much time?
// thinking to isntead make positions a string so i can regex match onto it

func allPositionsEndInZ(positions []string) bool {
	for _, position := range positions {
		if !strings.HasSuffix(position, "Z") {
			return false
		}
	}
	return true
}

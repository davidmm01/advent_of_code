package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type xToYMap struct {
	destinationRangeStart int
	sourceRangeStart      int
	rangeLength           int
}

var keys = []string{
	"seed,soil",
	"soil,fertilizer",
	"fertilizer,water",
	"water,light",
	"light,temperature",
	"temperature,humidity",
	"humidity,location",
}

var xToYRegex = regexp.MustCompile(`^([a-z]*)-to-([a-z]*) map:$`)

func getDestination(dataSets []xToYMap, source int) int {
	destination := -1

	for _, dataSet := range dataSets {
		if source >= dataSet.sourceRangeStart && source <= dataSet.sourceRangeStart+dataSet.rangeLength {
			// fmt.Println("      its a match!")
			// if this holds, then it fell within the range, find the corresponding range number\
			// fmt.Println("      source:", source)
			delta := source - dataSet.sourceRangeStart
			// fmt.Println("      delta:", delta)
			destination = dataSet.destinationRangeStart + delta
			// fmt.Println("      destination:", destination)

			if destination != -1 {
				return destination
			}

		}
	}

	return source
}

func main() {
	seedsToPlant, xToYMaps := processInput()

	var locations = []int{}
	var source int

	for _, seed := range seedsToPlant {
		// fmt.Println("for seed:", seed)
		source = seed
		for _, key := range keys {
			// fmt.Println("  for key:", key)
			// fmt.Println("    source:", source)
			source = getDestination(xToYMaps[key], source)
			// fmt.Println("    destination:", source)

		}
		locations = append(locations, source)
	}

	fmt.Println("Part 1:", slices.Min(locations))
}

func processInput() ([]int, map[string][]xToYMap) {
	in, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal("couldn't even read the file dave wtf")
	}

	lines := strings.Split(string(in), "\n")

	seedsToPlantString := lines[0]
	seedsToPlantString, ok := strings.CutPrefix(seedsToPlantString, "seeds: ")
	if !ok {
		log.Fatal("bad input puzzle")
	}
	seedsToPlantStringSlice := strings.Split(seedsToPlantString, " ")

	seedsToPlant := []int{}
	for _, seed := range seedsToPlantStringSlice {
		seedAsInt, err := strconv.Atoi(seed)
		if err != nil {
			log.Fatal("bad input puzzle")
		}
		seedsToPlant = append(seedsToPlant, seedAsInt)
	}

	// fmt.Println("seedsToPlant:", seedsToPlant)

	// chop off the first 2 lines - we are done with them
	lines = lines[2:]

	var currentFrom string
	var currentTo string
	var currentKey string

	// xToYMaps has keys of form `currentFrom,currentTo`
	xToYMaps := make(map[string][]xToYMap)

	for _, line := range lines {
		// skip divider lines
		if line == "" {
			continue
		}

		// if the line matches the regex, we know we are entering a new X to Y section
		matches := xToYRegex.FindAllStringSubmatch(line, -1)
		if len(matches) > 0 {
			currentFrom = matches[0][1]
			currentTo = matches[0][2]
			currentKey = fmt.Sprintf("%s,%s", currentFrom, currentTo)
			xToYMaps[currentKey] = []xToYMap{}
		} else {
			// else we are looking at a set of ranges, add this to the collection
			stringNumberSlice := strings.Split(line, " ")
			destinationRangeStart, _ := strconv.Atoi(stringNumberSlice[0])
			sourceRangeStart, _ := strconv.Atoi(stringNumberSlice[1])
			rangeLength, _ := strconv.Atoi(stringNumberSlice[2])

			currentMap := xToYMap{
				destinationRangeStart: destinationRangeStart,
				sourceRangeStart:      sourceRangeStart,
				rangeLength:           rangeLength,
			}

			xToYMaps[currentKey] = append(xToYMaps[currentKey], currentMap)
		}
	}

	return seedsToPlant, xToYMaps
}

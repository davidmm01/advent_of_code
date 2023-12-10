package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// gonna need this in lots of places
var xCoordMax int
var yCoordMax int

func main() {
	in, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal("couldn't even read the file dave wtf")
	}

	lines := strings.Split(string(in), "\n")

	yCoordMax = len(lines) - 1
	xCoordMax = len(lines[0]) - 1

	// if the next thing we saw wasn't a number, OR we reached the end of the line...
	// if we are currently scanning through a number then this must be the end of it
	numbers := getNumbers(lines)

	sum := 0
	for _, number := range numbers {
		if hasAdjacentSymbol(lines, number) {
			sum += number.value
		}
	}

	fmt.Println("Part 1 - part number sum:", sum)

	gearMatrix := getGearMatrix(lines)
	for _, number := range numbers {
		adjacentGears := getAdjacentGears(lines, number)
		for _, adjacentGear := range adjacentGears {
			gearMatrix[adjacentGear.yCoord][adjacentGear.xCoord].adjacentNumbers = append(gearMatrix[adjacentGear.yCoord][adjacentGear.xCoord].adjacentNumbers, number)
		}
	}

	gearRatioSum := 0

	for y := 0; y <= yCoordMax; y++ {
		for x := 0; x <= xCoordMax; x++ {
			if gearMatrix[y][x] != nil {
				if len(gearMatrix[y][x].adjacentNumbers) == 2 {
					gearRatio := gearMatrix[y][x].adjacentNumbers[0].value * gearMatrix[y][x].adjacentNumbers[1].value
					gearRatioSum += gearRatio
				}
			}
		}
	}

	fmt.Println("Part 2 - gear ratio sum:", gearRatioSum)
}

type Gear struct {
	xCoord          int
	yCoord          int
	adjacentNumbers []Number
}

func getGearMatrix(lines []string) [][]*Gear {
	gears := make([][]*Gear, len(lines))
	for i := range gears {
		gears[i] = make([]*Gear, len(lines[0]))
	}

	for y := 0; y <= yCoordMax; y++ {
		for x := 0; x <= xCoordMax; x++ {
			if string(lines[y][x]) == "*" {
				gears[y][x] = &Gear{
					xCoord:          x,
					yCoord:          y,
					adjacentNumbers: []Number{},
				}
			}
		}
	}

	return gears
}

func getAdjacentGears(lines []string, number Number) []*Gear {
	perimeterCoords := getPerimeterCoords(lines, number)

	var adjacentGears []*Gear

	for _, coord := range perimeterCoords {
		char := string(lines[coord.y][coord.x])
		if char == "*" {
			adjacentGears = append(adjacentGears, &Gear{
				xCoord:          coord.x,
				yCoord:          coord.y,
				adjacentNumbers: []Number{},
			})
		}
	}

	return adjacentGears
}

func hasAdjacentSymbol(lines []string, number Number) bool {
	perimeterCoords := getPerimeterCoords(lines, number)
	for _, coord := range perimeterCoords {
		char := string(lines[coord.y][coord.x])
		if _, err := strconv.Atoi(char); err != nil {
			if char != "." {
				return true
			}
		}
	}
	return false
}

type CoordPair struct {
	x int
	y int
}

func getPerimeterCoords(lines []string, number Number) (coordPairs []CoordPair) {
	for y := number.yCoord - 1; y <= number.yCoord+1; y++ {
		for x := number.startingXCoord - 1; x <= number.finishingXCoord; x++ {
			coord := CoordPair{x: x, y: y}
			if coordIsPossible(coord) && coordNotInNumber(coord, number) {
				coordPairs = append(coordPairs, coord)
			}
		}
	}
	return coordPairs
}

func coordIsPossible(coord CoordPair) bool {
	return !(coord.x < 0 || coord.x > xCoordMax || coord.y < 0 || coord.y > yCoordMax)
}

func coordNotInNumber(coord CoordPair, number Number) bool {
	if (number.startingXCoord <= coord.x && coord.x < number.finishingXCoord) && coord.y == number.yCoord {
		return false
	}
	return true
}

type Number struct {
	value           int
	startingXCoord  int
	finishingXCoord int
	yCoord          int
}

func getNumbers(lines []string) []Number {
	var numbers []Number

	for y := 0; y <= yCoordMax; y++ {
		startingX := 0
		traversingNumber := false

		for x := 0; x <= xCoordMax; x++ {
			_, err := strconv.Atoi(string(lines[y][x]))
			isNumber := err == nil

			if !isNumber {
				// if what we are currently looking at is not a number and we were traversing a number, then
				// we reached the end of the number, it should be added to the list of numbers
				if traversingNumber {
					value, err := strconv.Atoi(lines[y][startingX:x])
					if err != nil {
						log.Fatal("your logic is wrong david")
					}

					numbers = append(numbers, Number{
						value:           value,
						startingXCoord:  startingX,
						finishingXCoord: x,
						yCoord:          y,
					})
				}
				startingX = 0
				traversingNumber = false
				continue
			}

			// special case handling, final character in the line is a number
			if isNumber && x == xCoordMax {
				// we must be traversing - if starting x is 0 than this must be a 1 digit number
				if startingX == 0 {
					startingX = x
				}

				value, err := strconv.Atoi(lines[y][startingX : x+1]) // add 1 since we didn't already count past as per previous case
				if err != nil {
					log.Fatal("your logic is wrong david")
				}

				numbers = append(numbers, Number{
					value:           value,
					startingXCoord:  startingX,
					finishingXCoord: x + 1, // add 1 since we didn't already count past as per previous case
					yCoord:          y,
				})

				startingX = 0
				traversingNumber = false
				continue
			}

			// what we saw is a number, so lets register we are traversing
			if !traversingNumber {
				traversingNumber = true
				startingX = x
			}

		}
	}
	return numbers
}

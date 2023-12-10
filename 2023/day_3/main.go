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
			fmt.Printf("number %d has nearby symbol\n", number.value)
			sum += number.value
		} else {

			fmt.Printf("number %d DOES NOT have nearby symbol\n", number.value)
		}

	}

	fmt.Println("Part 1 - sum:", sum)
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

			if err != nil || x == xCoordMax {

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

			if !traversingNumber {
				traversingNumber = true
				startingX = x
			}

		}
	}
	return numbers
}

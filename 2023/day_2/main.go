package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func part1() {
	in, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal("couldn't even read the file dave wtf")
	}

	lines := strings.Split(string(in), "\n")

	sum := 0

	for _, line := range lines {
		id, err := getGameId(line)
		if err != nil {
			log.Fatal("Can't get game id from input line")
		}

		if areGamesPossible(line) {
			sum += id
		}
	}

	fmt.Println("Part 1 - sum:", sum)
}

func getGameId(line string) (int, error) {
	gameId := strings.Split(line, ":")[0]
	r, err := regexp.Compile(`Game ([\d]*)$`)
	if err != nil {
		return 0, err
	}
	m := r.FindStringSubmatch(gameId)
	return strconv.Atoi(m[1])
}

const (
	blueMax  = 14
	greenMax = 13
	redMax   = 12
)

var blueRegex = regexp.MustCompile(`([\d]*) blue`)
var greenRegex = regexp.MustCompile(`([\d]*) green`)
var redRegex = regexp.MustCompile(`([\d]*) red`)

func isGamePossible(game string, regex *regexp.Regexp, max int) bool {
	m := regex.FindStringSubmatch(game)
	if len(m) > 0 {
		asInt, err := strconv.Atoi(m[1])
		if err != nil {
			log.Fatalf("couldnt convert string to int: %s", m[1])
		}
		return asInt <= max
	}
	return true
}

func areGamesPossible(line string) bool {
	postColon := strings.Split(line, ":")[1]
	games := strings.Split(postColon, ";")
	for _, game := range games {
		if !isGamePossible(game, blueRegex, blueMax) ||
			!isGamePossible(game, greenRegex, greenMax) ||
			!isGamePossible(game, redRegex, redMax) {
			return false
		}
	}

	return true
}

func getColourRequired(game string, regex *regexp.Regexp) int {
	m := regex.FindStringSubmatch(game)
	if len(m) > 0 {
		asInt, err := strconv.Atoi(m[1])
		if err != nil {
			log.Fatalf("couldnt convert string to int: %s", m[1])
		}
		return asInt
	}
	return 0
}

func getRequired(line string) (int, int, int) {
	postColon := strings.Split(line, ":")[1]
	games := strings.Split(postColon, ";")

	blueMin := 0
	greenMin := 0
	redMin := 0

	for _, game := range games {
		blueRequired := getColourRequired(game, blueRegex)
		greenRequired := getColourRequired(game, greenRegex)
		redRequired := getColourRequired(game, redRegex)

		if blueRequired > blueMin {
			blueMin = blueRequired
		}
		if greenRequired > greenMin {
			greenMin = greenRequired
		}
		if redRequired > redMin {
			redMin = redRequired
		}
	}

	return blueMin, greenMin, redMin
}

func part2() {
	in, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal("couldn't even read the file dave wtf")
	}

	lines := strings.Split(string(in), "\n")

	sum := 0

	for _, line := range lines {
		blueMinRequired, greenMinRequired, redMinRequired := getRequired(line)
		cubePower := blueMinRequired * greenMinRequired * redMinRequired
		sum += cubePower
	}

	fmt.Println("Part 2 - sum:", sum)
}

func main() {
	part1()
	part2()
}

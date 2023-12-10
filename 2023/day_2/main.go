package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
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
		fmt.Println(id)

		if areGamesPossible(line) {
			fmt.Printf("Game %d was possible\n", id)
			sum += id
		} else {
			fmt.Printf("Game %d was not possbile\n", id)
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
		fmt.Println("game:", game)
		fmt.Println("asInt:", asInt)
		fmt.Println("max:", max)
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

package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Card struct {
	Id             string
	WinningNumbers map[int]bool
	PlayingNumbers []int
	Copies         int
}

func main() {
	in, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal("couldn't even read the file dave wtf")
	}

	lines := strings.Split(string(in), "\n")

	cards := getCards(lines)

	part1Output(cards)
	part2Output(cards)
}

func part2Output(cards []Card) {
	// build up all the copies
	for cardIndex, card := range cards {
		matches := countCardMatchingNumbers(card)
		if matches > 0 {
			for copy := 1; copy <= card.Copies; copy++ {
				for winnerIndex := cardIndex + 1; winnerIndex <= cardIndex+matches; winnerIndex++ {
					cards[winnerIndex].Copies++
				}
			}
		}
	}
	// iterate through and count up all the copies
	totalCards := 0
	for _, card := range cards {
		totalCards += card.Copies
	}
	fmt.Println("Part 2 - total cards:", totalCards)
}

func countCardMatchingNumbers(card Card) int {
	count := 0
	for _, playingNumber := range card.PlayingNumbers {
		if card.WinningNumbers[playingNumber] {
			count++
		}
	}
	return count
}

func part1Output(cards []Card) {
	totalPoints := 0
	for _, card := range cards {
		cardPoints := 0
		for _, playingNumber := range card.PlayingNumbers {
			if card.WinningNumbers[playingNumber] {
				if cardPoints == 0 {
					cardPoints = 1
				} else {
					cardPoints *= 2
				}
			}
		}
		totalPoints += cardPoints
	}
	fmt.Println("Part 1 - total points:", totalPoints)
}

func getCards(lines []string) (cards []Card) {
	idRegexString := `Card[\s]*([\d]*)`
	idRegex := regexp.MustCompile(idRegexString)

	numberRegexString := `[\s]*([\d]*)[\s]*`
	numberRegex := regexp.MustCompile(numberRegexString)

	for _, line := range lines {
		preColon := strings.Split(line, ":")[0]
		id := idRegex.FindStringSubmatch(preColon)[1]

		postColon := strings.Split(line, ":")[1]
		winningNumbersRaw := strings.Split(postColon, "|")[0]

		var winningNumbersMap = make(map[int]bool)
		for _, winningNumberDirty := range numberRegex.FindAllString(winningNumbersRaw, -1) {
			stripped := strings.Replace(winningNumberDirty, " ", "", -1)
			// converting it to a number is a good way to check we clean up the data nice
			value, err := strconv.Atoi(stripped)
			if err != nil {
				log.Fatal("Code is wrong")
			}
			winningNumbersMap[value] = true
		}

		playingNumbersRaw := strings.Split(postColon, "|")[1]
		playingNumbers := []int{}
		for _, playingNumbersDirty := range numberRegex.FindAllString(playingNumbersRaw, -1) {
			stripped := strings.Replace(playingNumbersDirty, " ", "", -1)
			// converting it to a number is a good way to check we clean up the data nice
			value, err := strconv.Atoi(stripped)
			if err != nil {
				log.Fatal("Code is wrong")
			}
			playingNumbers = append(playingNumbers, value)
		}

		cards = append(cards, Card{
			Id:             id,
			WinningNumbers: winningNumbersMap,
			PlayingNumbers: playingNumbers,
			Copies:         1,
		})

	}
	return cards
}

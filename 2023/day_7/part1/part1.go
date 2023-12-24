package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type CardStrength map[string]int

type HandTypeStrength map[HandType]int

type HandType int

var cardStrengthMap = makeCardStrengthMap()
var handStrengthMap = makeHandStrengthMap()

const (
	fiveKind HandType = iota
	fourKind
	fullHouse
	threeKind
	twoPair
	onePair
	highCard
)

type Game struct {
	hand             string
	bid              int
	handType         HandType
	handTypeStrength int
}

func main() {
	in, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal("couldn't even read the file dave wtf")
	}

	lines := strings.Split(string(in), "\n")
	var games []Game

	for _, line := range lines {
		data := strings.Split(line, " ")

		hand := data[0]
		bid, _ := strconv.Atoi(data[1])
		handType, err := determineHandType(hand)
		if err != nil {
			log.Fatal(err)
		}
		handTypeStrength := handStrengthMap[handType]

		game := Game{
			hand:             hand,
			bid:              bid,
			handType:         handType,
			handTypeStrength: handTypeStrength,
		}

		games = append(games, game)
	}

	sort.Slice(games, func(i, j int) bool {
		// if hand types are equal, let leading card strength dictate the winner
		if games[i].handTypeStrength == games[j].handTypeStrength {
			return leadingCardStrengthCompare(games[i], games[j])
		}
		return games[i].handTypeStrength < games[j].handTypeStrength
	})

	totalWinnings := 0
	for i, game := range games {
		totalWinnings += game.bid * (i + 1)
	}

	fmt.Println("Part 1:", totalWinnings)
}

func makeCardStrengthMap() CardStrength {
	cardStrengthMap := make(CardStrength)
	cardStrengthMap["2"] = 1
	cardStrengthMap["3"] = 2
	cardStrengthMap["4"] = 3
	cardStrengthMap["5"] = 4
	cardStrengthMap["6"] = 5
	cardStrengthMap["7"] = 6
	cardStrengthMap["8"] = 7
	cardStrengthMap["9"] = 8
	cardStrengthMap["T"] = 9
	cardStrengthMap["J"] = 10
	cardStrengthMap["Q"] = 11
	cardStrengthMap["K"] = 12
	cardStrengthMap["A"] = 13
	return cardStrengthMap
}

func makeHandStrengthMap() HandTypeStrength {
	handTypeStrengthMap := make(HandTypeStrength)
	handTypeStrengthMap[highCard] = 1
	handTypeStrengthMap[onePair] = 2
	handTypeStrengthMap[twoPair] = 3
	handTypeStrengthMap[threeKind] = 4
	handTypeStrengthMap[fullHouse] = 5
	handTypeStrengthMap[fourKind] = 6
	handTypeStrengthMap[fiveKind] = 7
	return handTypeStrengthMap
}

func determineHandType(hand string) (HandType, error) {
	seenChars := make(map[rune]int)
	for _, char := range hand {
		seenChars[char]++
	}

	highest := 0
	seen3 := false
	seen2 := 0
	for _, value := range seenChars {
		// we care about:
		// highest number of matching cards
		if value > highest {
			highest = value
		}

		// how many pairs there were (for full house and 2 pair)
		if value == 2 {
			seen2++
		}

		// if there was a 3 of a kind (for full house)
		if value == 3 {
			seen3 = true
		}
	}

	switch highest {
	case 5:
		return fiveKind, nil
	case 4:
		return fourKind, nil
	case 3:
		if seen3 && seen2 == 1 {
			return fullHouse, nil
		}
		return threeKind, nil
	case 2:
		if seen2 == 2 {
			return twoPair, nil
		}
		return onePair, nil
	case 1:
		return highCard, nil
	}

	return -1, errors.New("logic is wrong, this error should not happen")
}

// leadingCardStrengthCompare answers "is game1 leading card strength less than game2"?
func leadingCardStrengthCompare(game1 Game, game2 Game) bool {
	for i := 0; i < len(game1.hand); i++ {
		game1Card := string(game1.hand[i])
		game2Card := string(game2.hand[i])
		if game1Card == game2Card {
			continue
		}
		return cardStrengthMap[game1Card] < cardStrengthMap[game2Card]
	}
	log.Fatal("hmm")
	return false
}

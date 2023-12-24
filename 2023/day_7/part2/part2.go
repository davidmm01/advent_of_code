package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type CardStrength map[string]int

type HandTypeStrength map[HandType]int

type HandType string

var cardStrengthMap = makeCardStrengthMap()
var handStrengthMap = makeHandStrengthMap()

const (
	fiveKind     HandType = "fivekind"
	fourKind              = "fourKind"
	fullHouse             = "fullHouse"
	threeKind             = "threeKind"
	twoPair               = "twoPair"
	onePair               = "onePair"
	highCard              = "highCard"
	handOfJokers          = "handOfJokers"
)

type Game struct {
	hand                         string
	jokerlessHand                string
	jokerlessHandType            HandType
	bestPossibleHandType         HandType
	jokers                       int
	bid                          int
	bestPossibleHandTypeStrength int
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
		jokerlessHand := strings.ReplaceAll(data[0], "J", "")
		bid, _ := strconv.Atoi(data[1])
		jokerlessHandType := determineHandType(jokerlessHand)

		jokers := strings.Count(hand, "J")

		bestPossibleHandType := determineBestPossibleHandType(hand, jokerlessHand, jokers, jokerlessHandType)
		bestPossibleHandTypeStrength := handStrengthMap[bestPossibleHandType]

		game := Game{
			hand:                         hand,
			jokerlessHand:                jokerlessHand,
			jokerlessHandType:            jokerlessHandType,
			bestPossibleHandType:         bestPossibleHandType,
			bid:                          bid,
			jokers:                       jokers,
			bestPossibleHandTypeStrength: bestPossibleHandTypeStrength,
		}

		games = append(games, game)
	}

	sort.Slice(games, func(i, j int) bool {
		// if hand types are equal, let leading card strength dictate the winner
		if games[i].bestPossibleHandTypeStrength == games[j].bestPossibleHandTypeStrength {
			return leadingCardStrengthCompare(games[i], games[j])
		}
		return games[i].bestPossibleHandTypeStrength < games[j].bestPossibleHandTypeStrength
	})

	totalWinnings := 0
	for i, game := range games {
		totalWinnings += game.bid * (i + 1)
	}

	fmt.Println("Part 2:", totalWinnings)
}

func makeCardStrengthMap() CardStrength {
	cardStrengthMap := make(CardStrength)
	cardStrengthMap["J"] = 1
	cardStrengthMap["2"] = 2
	cardStrengthMap["3"] = 3
	cardStrengthMap["4"] = 4
	cardStrengthMap["5"] = 5
	cardStrengthMap["6"] = 6
	cardStrengthMap["7"] = 7
	cardStrengthMap["8"] = 8
	cardStrengthMap["9"] = 9
	cardStrengthMap["T"] = 10
	cardStrengthMap["Q"] = 11
	cardStrengthMap["K"] = 12
	cardStrengthMap["A"] = 13
	return cardStrengthMap
}

func makeHandStrengthMap() HandTypeStrength {
	handTypeStrengthMap := make(HandTypeStrength)
	handTypeStrengthMap[handOfJokers] = 0
	handTypeStrengthMap[highCard] = 1
	handTypeStrengthMap[onePair] = 2
	handTypeStrengthMap[twoPair] = 3
	handTypeStrengthMap[threeKind] = 4
	handTypeStrengthMap[fullHouse] = 5
	handTypeStrengthMap[fourKind] = 6
	handTypeStrengthMap[fiveKind] = 7
	return handTypeStrengthMap
}

func determineHandType(hand string) HandType {
	seenChars := make(map[rune]int)
	for _, char := range hand {
		seenChars[char]++
	}

	return getHandTypeFromCharCount(seenChars)
}

func getHandTypeFromCharCount(seenChars map[rune]int) HandType {
	// we care about:
	// highest number of matching cards
	// how many pairs there were (for full house and 2 pair)
	// if there was a 3 of a kind (for full house)
	highest := 0
	seen3 := false
	seen2 := 0
	for _, value := range seenChars {

		if value > highest {
			highest = value
		}

		if value == 2 {
			seen2++
		}

		if value == 3 {
			seen3 = true
		}
	}

	switch highest {
	case 5:
		return fiveKind
	case 4:
		return fourKind
	case 3:
		if seen3 && seen2 == 1 {
			return fullHouse
		}
		return threeKind
	case 2:
		if seen2 == 2 {
			return twoPair
		}
		return onePair
	case 1:
		return highCard
	case 0:
		return handOfJokers
	}

	panic(fmt.Sprintf("Blowing up. highest=%d, seen3=%v, seen2=%d, seenChars=%+v\n", highest, seen3, seen2, seenChars))
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

func determineBestPossibleHandType(hand string, jokerlessHand string, jokers int, jokerlessHandType HandType) HandType {
	if jokers == 0 {
		return determineHandType(hand)
	}

	seenChars := make(map[rune]int)
	for _, char := range jokerlessHand {
		seenChars[char]++
	}

	// just make the already-best card better by jokers, this is always the strongest move since the char type wont contribute to leading card strength comparisons
	var bestChar rune
	mostSeen := 0
	for char, timesSeen := range seenChars {
		if timesSeen > mostSeen {
			mostSeen = timesSeen
			bestChar = char
		}
	}

	seenChars[bestChar] += jokers
	return getHandTypeFromCharCount(seenChars)
}

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

var getNumbersFromInputRegex = regexp.MustCompile(`[\d]*`)

func main() {
	in, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal("couldn't even read the file dave wtf")
	}

	lines := strings.Split(string(in), "\n")

	for _, line := range lines {
		historyStrings := strings.Split(line, " ")
		history := []int{}
		for _, historyString := range historyStrings {
			asInt, _ := strconv.Atoi(historyString)
			history = append(history, asInt)
		}
		fmt.Println("** starting next history:", history)
		fmt.Println(progressHistory(history))
	}
}

func historyProgressor(history []int) {
	// this thing will start up the recursion
}

func progressHistory(history []int) int {
	done := true
	if All()


	nextHistory := []int{}
	for i := 0; i < len(history)-1; i++ {
		dataPoint := history[i+1] - history[i]
		nextHistory = append(nextHistory, dataPoint)
		if dataPoint != 0 {
			done = false
		}
	}
	fmt.Println(nextHistory)
	if !done {
		bubbleUp := progressHistory(nextHistory)
		fmt.Printf("line %v has bubble up of %d\n", history, bubbleUp+history[len(history)-1])
		return bubbleUp + history[len(history)-1]
	} else {
		return 0
	}
}

func allZero(history []int) bool {
	for _, value := range history {
		if value != 0 {
			return false
		}
	}
	return true
}

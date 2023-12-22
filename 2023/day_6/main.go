package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var numberRegex = regexp.MustCompile(`(\d)+`)

type Race struct {
	time     int
	distance int
}

func main() {
	part1()
	part2()
}

func part1() {
	in, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal("couldn't even read the file dave wtf")
	}

	lines := strings.Split(string(in), "\n")
	timesRaw := numberRegex.FindAllStringSubmatch(lines[0], -1)
	distancesRaw := numberRegex.FindAllStringSubmatch(lines[1], -1)
	if len(timesRaw) != len(distancesRaw) {
		log.Fatal("broooo")
	}

	var races []Race

	for i := 0; i < len(timesRaw); i++ {
		time, _ := strconv.Atoi(timesRaw[i][0])
		distance, _ := strconv.Atoi(distancesRaw[i][0])
		race := Race{
			time:     time,
			distance: distance,
		}
		races = append(races, race)
	}

	// fmt.Println(races)
	var waysToWin []int

	for _, race := range races {
		waysToWinRace := 0

		for chargeTime := 1; chargeTime <= race.time-1; chargeTime++ { // dont bother with cases chargeTime=0 or chargeTime=time, as these will always lead to distance == 0
			timeDriving := race.time - chargeTime

			distance := chargeTime * timeDriving

			if distance > race.distance {
				waysToWinRace += 1
			}
		}
		waysToWin = append(waysToWin, waysToWinRace)
	}

	answer := 1
	for _, i := range waysToWin {
		answer *= i
	}
	fmt.Println("Part 1:", answer)
}

func part2() {
}

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	in, err := os.ReadFile("input.txt")
	if err != nil {
		return
	}

	lines := strings.Split(string(in), "\n")

	for _, line := range lines {
		fmt.Println(line)

		// get game id, split by :

		// get
	}

}

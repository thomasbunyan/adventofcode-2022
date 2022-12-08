package day2

import (
	"fmt"
	"strings"

	"github.com/thomasbunyan/adventofcode-2022/common"
)

const (
	WIN  = 6
	LOSS = 0
	DRAW = 3
)

type ShapeData struct {
	Alias string
	Value int32
	Win   string
	Loose string
}

var Guide = map[string]ShapeData{
	"A": {Alias: "rock", Value: 1, Win: "Z", Loose: "Y"},
	"B": {Alias: "paper", Value: 2, Win: "X", Loose: "Z"},
	"C": {Alias: "scissors", Value: 3, Win: "Y", Loose: "X"},
	"X": {Alias: "rock", Value: 1, Win: "C"},
	"Y": {Alias: "paper", Value: 2, Win: "A"},
	"Z": {Alias: "scissors", Value: 3, Win: "B"},
}

func Run() {
	lines, err := common.Read("./day2/input.txt")
	if err != nil {
		panic(err)
	}

	scoreA := 0
	scoreB := 0

	for _, line := range lines {
		input := strings.Split(line, " ")
		opponent := input[0]
		player := input[1]

		scoreA += int(Guide[player].Value)

		if Guide[opponent].Win == player {
			// opponent wins
			scoreA += 0
		} else if Guide[player].Win == opponent {
			// player wins
			scoreA += 6
		} else {
			// draw
			scoreA += 3
		}

		switch player {
		case "X":
			// loose
			scoreB += int(Guide[Guide[opponent].Win].Value)
			scoreB += 0
		case "Y":
			// draw
			scoreB += int(Guide[opponent].Value)
			scoreB += 3
		case "Z":
			// win
			scoreB += int(Guide[Guide[opponent].Loose].Value)
			scoreB += 6
		}
	}

	fmt.Printf("Score A: %v\n", scoreA)
	fmt.Printf("Score B: %v\n", scoreB)
}

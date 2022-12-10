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
	Lose  string
}

var Guide = map[string]ShapeData{
	"A": {Alias: "rock", Value: 1, Win: "Z", Lose: "Y"},
	"B": {Alias: "paper", Value: 2, Win: "X", Lose: "Z"},
	"C": {Alias: "scissors", Value: 3, Win: "Y", Lose: "X"},
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
			scoreA += LOSS
		} else if Guide[player].Win == opponent {
			// player wins
			scoreA += WIN
		} else {
			// draw
			scoreA += DRAW
		}

		switch player {
		case "X":
			// lose
			scoreB += int(Guide[Guide[opponent].Win].Value)
			scoreB += LOSS
		case "Y":
			// draw
			scoreB += int(Guide[opponent].Value)
			scoreB += DRAW
		case "Z":
			// win
			scoreB += int(Guide[Guide[opponent].Lose].Value)
			scoreB += WIN
		}
	}

	fmt.Printf("Score A: %v\n", scoreA)
	fmt.Printf("Score B: %v\n", scoreB)
}

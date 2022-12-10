package day6

import (
	"fmt"

	"github.com/thomasbunyan/adventofcode-2022/common"
)

func Run() {
	lines, err := common.Read("./day6/input.txt")
	if err != nil {
		panic(err)
	}

	line := lines[0]

	const (
		packetOffset  = 4
		messageOffset = 14
	)

	var packetMarker int
	for i := 0; i < len(line)-packetOffset; i++ {
		window := line[i : i+packetOffset]
		if AllUnique(window) {
			packetMarker = i + packetOffset
			break
		}
	}

	var messageMarker int
	for i := 0; i < len(line)-messageOffset; i++ {
		window := line[i : i+messageOffset]
		if AllUnique(window) {
			messageMarker = i + messageOffset
			break
		}
	}

	fmt.Println(packetMarker)
	fmt.Println(messageMarker)
}

func AllUnique(s string) bool {
	seen := make(map[rune]bool)
	for _, c := range s {
		if seen[c] {
			return false
		} else {
			seen[c] = true
		}
	}
	return true
}

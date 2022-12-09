package day3

import (
	"fmt"
	"strings"

	"github.com/thomasbunyan/adventofcode-2022/common"
)

const ABC = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const abc = "abcdefghijklmnopqrstuvwxyz"

func getValue(input string) int {
	index := strings.Index(abc, input) + 1
	if index == 0 {
		index = strings.Index(ABC, input) + 27
	}
	return index
}

// Part 1
func Run() {
	lines, err := common.Read("./day3/input.txt")
	if err != nil {
		panic(err)
	}

	var items []string

	for _, line := range lines {
		first := line[:len(line)/2]
		second := line[len(line)/2:]

		m := make(map[string]bool)

		for i := 0; i < len(first); i++ {
			m[first[i:i+1]] = true
		}

		for i := 0; i < len(first); i++ {
			if m[second[i:i+1]] {
				items = append(items, second[i:i+1])
				break
			}
		}
	}

	var score int
	for _, item := range items {
		score += getValue(item)
	}

	fmt.Printf("Score: %v\n", score)
}

// Part 2
func Run2() {
	lines, err := common.Read("./day3/input.txt")
	if err != nil {
		panic(err)
	}

	var items []string

	for i := 0; i < len(lines); i += 3 {
		first := lines[i]
		second := lines[i+1]
		third := lines[i+2]

		m := make(map[string]int)

		for i := 0; i < len(first); i++ {
			m[first[i:i+1]] = 1
		}

		for i := 0; i < len(second); i++ {
			if m[second[i:i+1]] == 1 {
				m[second[i:i+1]] += 1
			}
		}

		for i := 0; i < len(third); i++ {
			if m[third[i:i+1]] == 2 {
				items = append(items, third[i:i+1])
				break
			}
		}
	}

	var score int
	for _, item := range items {
		score += getValue(item)
	}

	fmt.Printf("Score: %v\n", score)
}

package day5

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/thomasbunyan/adventofcode-2022/common"
)

// part 1
func Run() {
	lines, err := common.Read("./day5/input.txt")
	if err != nil {
		panic(err)
	}

	var (
		numOfStacks  int
		state        []string
		instructions []string
	)

	for i, line := range lines {
		if line == "" {
			state = lines[:i-1]
			instructions = lines[i:]

			numOfStacks = len(strings.Split(lines[i-1], "   "))

			break
		}
	}

	stacks := make([][]string, numOfStacks)
	for _, row := range state {
		cols := SplitN(row, 3)
		for i, col := range cols {
			if col != "   " {
				stacks[i] = append(stacks[i], col)
			}
		}
	}

	for _, instruction := range instructions[1:] {
		re := regexp.MustCompile(`^move (?P<q>\d+) from (?P<s>\d+) to (?P<d>\d+)$`)
		match := re.FindStringSubmatch(instruction)
		quant, _ := strconv.Atoi(match[1])
		source, _ := strconv.Atoi(match[2])
		dest, _ := strconv.Atoi(match[3])

		pop := stacks[source-1][:quant]
		sort.SliceStable(pop, func(i, j int) bool {
			return i > j
		})

		stacks[source-1] = stacks[source-1][quant:]
		// something I don't understand here
		stacks[dest-1] = append(append([]string(nil), pop...), stacks[dest-1]...)
	}

	var tops string
	for _, stack := range stacks {
		top := stack[0]
		tops = tops + strings.Replace(strings.Replace(top, "]", "", -1), "[", "", -1)
	}

	fmt.Println(tops)
}

// part 2
func Run2() {
	lines, err := common.Read("./day5/input.txt")
	if err != nil {
		panic(err)
	}

	var (
		numOfStacks  int
		state        []string
		instructions []string
	)

	for i, line := range lines {
		if line == "" {
			state = lines[:i-1]
			instructions = lines[i:]

			numOfStacks = len(strings.Split(lines[i-1], "   "))

			break
		}
	}

	stacks := make([][]string, numOfStacks)
	for _, row := range state {
		cols := SplitN(row, 3)
		for i, col := range cols {
			if col != "   " {
				stacks[i] = append(stacks[i], col)
			}
		}
	}

	for _, instruction := range instructions[1:] {
		re := regexp.MustCompile(`^move (?P<q>\d+) from (?P<s>\d+) to (?P<d>\d+)$`)
		match := re.FindStringSubmatch(instruction)
		quant, _ := strconv.Atoi(match[1])
		source, _ := strconv.Atoi(match[2])
		dest, _ := strconv.Atoi(match[3])

		pop := stacks[source-1][:quant]

		stacks[source-1] = stacks[source-1][quant:]
		// something I don't understand here
		stacks[dest-1] = append(append([]string(nil), pop...), stacks[dest-1]...)
	}

	var tops string
	for _, stack := range stacks {
		top := stack[0]
		tops = tops + strings.Replace(strings.Replace(top, "]", "", -1), "[", "", -1)
	}

	fmt.Println(tops)
}

func SplitN(s string, n int) []string {
	var out []string
	for i := 0; i < len(s); i += n + 1 {
		var upper int
		if i+n > len(s) {
			upper = len(s)
		} else {
			upper = i + n
		}
		out = append(out, s[i:upper])
	}
	return out
}

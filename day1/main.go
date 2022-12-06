package day1

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/thomasbunyan/adventofcode-2022/common"
)

func Run() {
	lines, err := common.Read("./day1/input.txt")
	if err != nil {
		panic(err)
	}

	var elves []int
	calories := 0

	for i := 0; i < len(lines); i++ {
		if lines[i] == "" {
			elves = append(elves, calories)
			calories = 0
			continue
		}
		a, _ := strconv.Atoi(lines[i])
		calories += int(a)
	}
	elves = append(elves, calories)

	sort.Ints(elves[:])

	fmt.Printf("Elf: %v\n", elves[(len(elves)-1)])

	fmt.Printf("Top 3: %v\n", elves[(len(elves)-1)]+elves[(len(elves)-2)]+elves[(len(elves)-3)])
}

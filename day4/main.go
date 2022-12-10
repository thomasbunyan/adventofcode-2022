package day4

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/thomasbunyan/adventofcode-2022/common"
)

func Run() {
	lines, err := common.Read("./day4/input.txt")
	if err != nil {
		panic(err)
	}

	count1 := 0
	count2 := 0

	for _, line := range lines {
		pair := strings.Split(line, ",")

		rangeA := strings.Split(pair[0], "-")
		rangeB := strings.Split(pair[1], "-")

		a1, _ := strconv.Atoi(rangeA[0])
		a2, _ := strconv.Atoi(rangeA[1])
		b1, _ := strconv.Atoi(rangeB[0])
		b2, _ := strconv.Atoi(rangeB[1])

		if a1 >= b1 && a2 <= b2 {
			count1 += 1
		} else if b1 >= a1 && b2 <= a2 {
			count1 += 1
		}

		if a1 <= b1 && a2 >= b1 {
			count2 += 1
		} else if b1 <= a1 && b2 >= a1 {
			count2 += 1
		}
	}

	fmt.Printf("Count 1: %v\n", count1)
	fmt.Printf("Count 2: %v\n", count2)
}

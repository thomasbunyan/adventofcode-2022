package main

import (
	"fmt"
	"log"
	"os"

	"github.com/thomasbunyan/adventofcode-2022/day1"
	"github.com/thomasbunyan/adventofcode-2022/day2"
	"github.com/thomasbunyan/adventofcode-2022/day3"
	"github.com/thomasbunyan/adventofcode-2022/day4"
	"github.com/thomasbunyan/adventofcode-2022/day5"
	"github.com/thomasbunyan/adventofcode-2022/day6"
)

type Input struct {
	Day string `json:"day"`
}

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		log.Fatal("[ERROR] missing the day to run")
	}

	input := &Input{}
	input.Day = os.Args[1]

	fmt.Printf("[INFO] running day %v\n", input.Day)

	switch input.Day {
	case "1":
		day1.Run()
	case "2":
		day2.Run()
	case "3":
		day3.Run()
	case "4":
		day4.Run()
	case "5":
		day5.Run()
	case "6":
		day6.Run()
	default:
		log.Fatal("[ERROR] invalid day")
	}
}

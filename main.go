package main

import (
	"fmt"
	"log"
	"os"

	"github.com/thomasbunyan/adventofcode-2022/day1"
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
	default:
		log.Fatal("[ERROR] invalid day")
	}
}

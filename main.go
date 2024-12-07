package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/WrathOP/advent-of-code-2024/utils"

	"github.com/WrathOP/advent-of-code-2024/solutions/day01"
	"github.com/WrathOP/advent-of-code-2024/solutions/day02"
	"github.com/WrathOP/advent-of-code-2024/solutions/day03"
	"github.com/WrathOP/advent-of-code-2024/solutions/day04"
	"github.com/WrathOP/advent-of-code-2024/solutions/day05"
	"github.com/WrathOP/advent-of-code-2024/solutions/day06"
	"github.com/WrathOP/advent-of-code-2024/solutions/day07"
)

type Solution interface {
	Part1(file *os.File) any
	Part2(file *os.File) any
}

var dayMapping = map[int]Solution{
	1: day01.Solutions{},
	2: day02.Solutions{},
	3: day03.Solutions{},
	4: day04.Solutions{},
	5: day05.Solutions{},
	6: day06.Solutions{},
	7: day07.Solutions{},
}

func main() {
	dayFlag := flag.Int("day", 0, "day to run (1-25)")
	partFlag := flag.Int("part", 1, "part to run (1 or 2)")
	allFlag := flag.Bool("all", false, "run all days")
	flag.Parse()

	if *dayFlag == 0 && !*allFlag {
		fmt.Println("Please specify a day (-day N) or use -all")
		os.Exit(1)
	}

	if *partFlag != 1 && *partFlag != 2 {
		fmt.Println("Part must be 1 or 2")
		os.Exit(1)
	}

	if *allFlag {
		for day := range dayMapping {
			runDay(day, 1)
			runDay(day, 2)
		}
		return
	}

	runDay(*dayFlag, *partFlag)
}

func runDay(day, part int) {
	solution, exists := dayMapping[day]
	if !exists {
		fmt.Printf("Day %d not implemented\n", day)
		return
	}

	file, err := utils.OpenInput(day)
	utils.Check(err)
	defer file.Close()

	var result any
	startTime := time.Now()

	switch part {
	case 1:
		result = solution.Part1(file)
	case 2:
		result = solution.Part2(file)
	}

	elapsedTime := time.Since(startTime)

	fmt.Printf("Day %d, Part %d: %v (Time taken: %v)\n", day, part, result, elapsedTime)
}

package day07

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/WrathOP/advent-of-code-2024/utils"
)

type Solutions struct{}

func part1Helper(target int, input []int, res int) int {
	// Base case: if we've used all numbers and reached target
	if len(input) == 0 {
		if res == target {
			return target
		}
		return 0
	}

	// Try multiplication: current result * next number
	if mult := part1Helper(target, input[1:], res*input[0]); mult > 0 {
		return mult
	}

	// Try addition: current result + next number
	if add := part1Helper(target, input[1:], res+input[0]); add > 0 {
		return add
	}

	return 0
}

func part2helper(target int, input []int, res int) int {
	// Base case: if we've used all numbers and reached target
	if len(input) == 0 {
		if res == target {
			return target
		}
		return 0
	}

	// Try multiplication: current result * next number
	if mult := part2helper(target, input[1:], res*input[0]); mult > 0 {
		return mult
	}

	// Try addition: current result + next number
	if add := part2helper(target, input[1:], res+input[0]); add > 0 {
		return add
	}

	// Try concatenation: current result + next number
	if concat := part2helper(target, input[1:], utils.MustAtoi(fmt.Sprintf("%d%d", res, input[0]))); concat > 0 {
		return concat
	}

	return 0
}
func (Solutions) Part1(file *os.File) any {
	input := make(chan string)
	go utils.ReadFileByLineIntoChannel(file, input)
	var wg sync.WaitGroup
	res := atomic.Int64{}

	for line := range input {
		buffer := strings.Split(line, ":")
		input := make([]int, 0, len(buffer[1]))
		for _, v := range strings.Split(strings.TrimSpace(buffer[1]), " ") {
			input = append(input, utils.MustAtoi(v))
		}

		wg.Add(1)
		go func() {
			defer wg.Done()
			res.Add(int64(part1Helper(utils.MustAtoi(buffer[0]), input, 0)))
		}()

	}

	wg.Wait()

	return res.Load()
}

func (Solutions) Part2(file *os.File) any {
	input := make(chan string)
	go utils.ReadFileByLineIntoChannel(file, input)
	var wg sync.WaitGroup
	res := atomic.Int64{}

	for line := range input {
		buffer := strings.Split(line, ":")
		input := make([]int, 0, len(buffer[1]))
		for _, v := range strings.Split(strings.TrimSpace(buffer[1]), " ") {
			input = append(input, utils.MustAtoi(v))
		}

		wg.Add(1)
		go func() {
			defer wg.Done()
			res.Add(int64(part2helper(utils.MustAtoi(buffer[0]), input, 0)))
		}()

	}

	wg.Wait()

	return res.Load()
}

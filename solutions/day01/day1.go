package day01

import (
	"math"
	"os"
	"sort"
	"strings"

	"github.com/WrathOP/advent-of-code-2024/utils"
)

type Solutions struct{}

func (s Solutions) Part1(file *os.File) any {
	slice1, slice2 := []int{}, []int{}
	input := make(chan string)
	go utils.ReadFileIntoChannel(file, input)

	for line := range input {
		temp := strings.Split(line, "   ")
		slice1 = append(slice1, utils.MustAtoi(temp[0]))
		slice2 = append(slice2, utils.MustAtoi(temp[1]))
	}

	sort.Ints(slice1)
	sort.Ints(slice2)

	res := 0
	for i := range slice1 {
		res += int(math.Abs(float64(slice1[i] - slice2[i])))
	}

	return res
}

func (s Solutions) Part2(file *os.File) any {
	slice1, slice2 := []int{}, []int{}
	input := make(chan string)
	go utils.ReadFileIntoChannel(file, input)

	for line := range input {
		temp := strings.Split(line, "   ")
		slice1 = append(slice1, utils.MustAtoi(temp[0]))
		slice2 = append(slice2, utils.MustAtoi(temp[1]))
	}

	occurrences := make(map[int]int)

	for _, value := range slice2 {
		occurrences[value]++
	}

	res := 0

	for _, value := range slice1 {
		res += value * occurrences[value]
	}

	return res
}

package day02

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/WrathOP/advent-of-code-2024/utils"
)

type Solutions struct{}

func part1Helper(arr []int) bool {
	isDesc := arr[0] > arr[1]
	for i := 1; i < len(arr); i++ {
		a, b := arr[i-1], arr[i]
		if i <= len(arr)-1 && !(isDesc == (a > b) && (1 <= utils.Abs(a-b) && utils.Abs(a-b) <= 3)) {
			return false
		}
	}
	return true
}

func part2Helper(arr []int) bool {
    if part1Helper(arr) {
        return true
    }

    // Brute forcing that bitch , wait this got a fancy name
    // Yes nerds call it DP ig idk I am dumb
    for i := 0; i < len(arr); i++ {
        newArr := make([]int, 0, len(arr)-1)
        newArr = append(newArr, arr[:i]...)
        newArr = append(newArr, arr[i+1:]...)

        if part1Helper(newArr) {
            return true
        }
    }
    return false
}

func (s Solutions) Part1(file *os.File) any {
	input := make(chan string)
	go utils.ReadFileIntoChannel(file, input)

	res := 0
	var mu sync.Mutex
	wg := sync.WaitGroup{}

	for line := range input {
		wg.Add(1)
		strSlice := strings.Fields(line)
		intSlice := make([]int, 0, len(strSlice))
		for _, str := range strSlice {
			num, err := strconv.Atoi(str)
			if err != nil {
				fmt.Printf("Error converting string to integer: %v\n", err)
				continue
			}
			intSlice = append(intSlice, num)
		}

		go func(slice []int) {
			defer wg.Done()
			mu.Lock()
			if part1Helper(intSlice) {
				res++
			}
			mu.Unlock()
		}(intSlice)

	}

	wg.Wait()
	return res
}

func (s Solutions) Part2(file *os.File) any {
	input := make(chan string)
	go utils.ReadFileIntoChannel(file, input)

	res := 0
	var mu sync.Mutex
	wg := sync.WaitGroup{}

	for line := range input {
		wg.Add(1)
		strSlice := strings.Fields(line)
		intSlice := make([]int, 0, len(strSlice))
		for _, str := range strSlice {
			num, err := strconv.Atoi(str)
			if err != nil {
				fmt.Printf("Error converting string to integer: %v\n", err)
				continue
			}
			intSlice = append(intSlice, num)
		}

		go func(slice []int) {
			defer wg.Done()
			mu.Lock()
			if part2Helper(intSlice) {
				res++
			}
			mu.Unlock()
		}(intSlice)

	}

	wg.Wait()
	return res
}

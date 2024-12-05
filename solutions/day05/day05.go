package day05

import (
	"os"
	"slices"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/WrathOP/advent-of-code-2024/utils"
)

type Solutions struct{}

func getMiddleOfArray(arr *[]int) int {
	return (*arr)[len(*arr)/2]
}

func helperPart1(adjList *map[int][]int, arr *[]int) bool {
	visited := make(map[int]bool, len(*arr))

	for _, a := range *arr {
		visited[a] = true
		for _, b := range (*adjList)[a] {
			if visited[b] {
				return false
			}
		}
	}

	return true
}

func (s Solutions) Part1(file *os.File) any {
	input := make(chan string)
	adjacencyList := make(map[int][]int)
	var res atomic.Int64
	var wg sync.WaitGroup
	processingAdjList := true

	go utils.ReadFileByLineIntoChannel(file, input)

	for line := range input {
		if line == "" {
			processingAdjList = false
			continue
		}

		if processingAdjList {
			func() {
				defer func() {
					if r := recover(); r != nil {
					}
				}()
				li := strings.Split(line, "|")
				adjacencyList[utils.MustAtoi(li[0])] = append(adjacencyList[utils.MustAtoi(li[0])], utils.MustAtoi(li[1]))
			}()
		} else {
			wg.Add(1)
			li := strings.Split(line, ",")
			intList := make([]int, len(li))
			for i, l := range li {
				intList[i] = utils.MustAtoi(l)
			}
			go func(intList []int) {
				defer wg.Done()
				temp := helperPart1(&adjacencyList, &intList)
				if temp {
					res.Add(int64(getMiddleOfArray(&intList)))
				}
			}(intList)
		}
	}

	wg.Wait()
	return res.Load()
}

// This is basically a sort function which is sorting the array based on the adjacency list
// So this is basically a bubble sort which iterates to half the arr and swaps the elements based on
// if the right element's adjacency list contains the left element if yes then swap otherwise not
//
// You might ask why we are iterating to half the array because we only need the middle in the end to add
func helperPart2(adjList map[int][]int, arr []int) []int {
	mid := (len(arr) / 2) + 1
	for i := 0; i < mid; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			// Check if key exists and value is present in one go
			if vals, ok := adjList[arr[j+1]]; ok && slices.Contains(vals, arr[j]) {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func (s Solutions) Part2(file *os.File) any {
	input := make(chan string)
	adjacencyList := make(map[int][]int)
	var res atomic.Int64
	var wg sync.WaitGroup
	processingAdjList := true

	go utils.ReadFileByLineIntoChannel(file, input)

	for line := range input {
		if line == "" {
			processingAdjList = false
			continue
		}

		if processingAdjList {
			func() {
				defer func() {
					if r := recover(); r != nil {
					}
				}()
				li := strings.Split(line, "|")
				adjacencyList[utils.MustAtoi(li[0])] = append(adjacencyList[utils.MustAtoi(li[0])], utils.MustAtoi(li[1]))
			}()
		} else {
			wg.Add(1)
			li := strings.Split(line, ",")
			intList := make([]int, len(li))
			for i, l := range li {
				intList[i] = utils.MustAtoi(l)
			}

			go func(intList []int) {
				defer wg.Done()
				temp := helperPart1(&adjacencyList, &intList)
				if temp {
					return
				}
				intList = helperPart2(adjacencyList, intList)
				res.Add(int64(getMiddleOfArray(&intList)))
			}(intList)
		}
	}

	wg.Wait()
	return res.Load()
}

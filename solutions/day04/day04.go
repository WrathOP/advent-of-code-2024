package day04

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/WrathOP/advent-of-code-2024/utils"
)

type Solutions struct{}

func preProcessData(input chan string) [][]string {
	temp := [][]string{}
	for line := range input {
		temp = append(temp, strings.Split(line, ""))
	}
	return temp
}

func checkAllDirections(arr *[][]string, i, j int) int {
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}, {-1, -1}, {-1, 1}, {1, -1}, {1, 1}}

	tempRes := 0
	for _, dir := range directions {
		temp := "X"
		for k := 1; k < 4; k++ {
			if i+dir[0]*k < 0 || i+dir[0]*k >= len(*arr) || j+dir[1]*k < 0 || j+dir[1]*k >= len((*arr)[i]) {
				break
			}
			temp += (*arr)[i+dir[0]*k][j+dir[1]*k]
		}

		fmt.Println(temp)

		if temp == "XMAS" {
			tempRes++
		}
	}

	return tempRes
}

func (s Solutions) Part1(file *os.File) any {
	input := make(chan string)
	go utils.ReadFileByLineIntoChannel(file, input)
	// If we could optimise this
	arr := preProcessData(input)

	res := 0

	var wg sync.WaitGroup

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			if arr[i][j] == "X" {
				wg.Add(1)
				go func(i, j int) {
					defer wg.Done()
					res += checkAllDirections(&arr, i, j)
				}(i, j)
			}
		}
	}

	wg.Wait()

	return res
}

func checkAllDirectionsForX(arr *[][]string, i, j int) (temp int) {

	diaDireactions := [][][]int{{{1, 1}, {0, 0}, {-1, -1}}, {{1, -1}, {0, 0}, {-1, 1}}}
	buffer := map[int]string{
		0: "",
		1: "",
	}

	for k, directions := range diaDireactions {
		for _, dir := range directions {
			func() {
				defer func() {
					if r := recover(); r != nil {
						// Handle out of range error
					}
				}()
				buffer[k] += (*arr)[i+dir[0]][j+dir[1]]
			}()
		}
	}
	if (buffer[0] == "MAS" && buffer[1] == "MAS") || (buffer[0] == "SAM" && buffer[1] == "SAM") || (buffer[0] == "MAS" && buffer[1] == "SAM") || (buffer[1] == "MAS" && buffer[0] == "SAM") {
		temp = 1
	}

	return
}

func (s Solutions) Part2(file *os.File) any {
	input := make(chan string)
	go utils.ReadFileByLineIntoChannel(file, input)
	// If we could optimise this
	arr := preProcessData(input)

	res := 0

	var wg sync.WaitGroup

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			if arr[i][j] == "A" {
				wg.Add(1)
				go func(i, j int) {
					defer wg.Done()
					res += checkAllDirectionsForX(&arr, i, j)
				}(i, j)
			}
		}
	}

	wg.Wait()

	return res
}

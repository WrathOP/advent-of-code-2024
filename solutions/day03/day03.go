package day03

import (
	"os"
	"regexp"

	"github.com/WrathOP/advent-of-code-2024/utils"
)

type Solutions struct{}

func regexHelperPart1(inputChannel chan string, outChannel chan []int) {
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)

	for input := range inputChannel {
		for _, match := range re.FindAllStringSubmatch(input, -1) {
			outChannel <- []int{utils.MustAtoi(match[1]), utils.MustAtoi(match[2])}
		}
	}

	close(outChannel)
}

func regexHelperPart2(inputChannel chan string, outChannel chan []int) {
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)|do\(\)|don't\(\)`)
    enabled := true

    for input := range inputChannel {
        matches := re.FindAllStringSubmatch(input, -1)
        for _, match := range matches {
            switch match[0] {
            case "do()":
                enabled = true
            case "don't()":
                enabled = false
            default:
                // This means it's a mul() instruction since it's the only other pattern
                if enabled {
                    num1 := utils.MustAtoi(match[1])
                    num2 := utils.MustAtoi(match[2])
                    outChannel <- []int{num1, num2}
                }
            }
        }
    }
    close(outChannel)
}

func (s Solutions) Part1(file *os.File) any {
	input := make(chan string)
	output := make(chan []int)

	res := 0

	go utils.ReadFileByLineIntoChannel(file, input)
	go regexHelperPart1(input, output)

	for arr := range output {
		res += arr[0] * arr[1]
	}

	return res
}

func (s Solutions) Part2(file *os.File) any {
	input := make(chan string)
	output := make(chan []int)

	res := 0
	go utils.ReadFileByLineIntoChannel(file, input)
	go regexHelperPart2(input, output)

	for arr := range output {
		// fmt.Println(arr)
		res += arr[0] * arr[1]
	}

	return res
}

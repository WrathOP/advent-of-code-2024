package utils

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadFileByLineIntoChannel(file *os.File, channel chan string) {
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		channel <- scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	close(channel)
}

func ReadFileIntoChannel(file *os.File, channel chan string) {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		channel <- scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	close(channel)
}

func MustAtoi(s string) int {
	i, err := strconv.Atoi(s)
	Check(err)
	return i
}

func OpenInput(day int) (*os.File, error) {
	path := filepath.Join("solutions", fmt.Sprintf("day%02d", day), "input.txt")
	return os.Open(path)
}

package day06

import (
	"os"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/WrathOP/advent-of-code-2024/utils"
)

type Solutions struct{}

type Point struct {
	x int
	y int
}

type Direction int

const (
	Up Direction = iota
	Right
	Down
	Left
)

type Soldier struct {
	location  Point
	direction Direction
}

var directionVectors = [4]Point{
	{-1, 0}, // Up
	{0, 1},  // Right
	{1, 0},  // Down
	{0, -1}, // Left
}

var turnMap = [4]Direction{
	Right, // Up turns Right
	Down,  // Right turns Down
	Left,  // Down turns Left
	Up,    // Left turns Up
}

func convert(input chan string) ([][]string, Soldier) {
	grid := make([][]string, 0)
	var x, y int
	i := 0
	for line := range input {
		temp := strings.Split(line, "")
		if pos := strings.Index(line, "^"); pos != -1 {
			x, y = i, pos
		}
		grid = append(grid, temp)
		i++
	}
	return grid, Soldier{location: Point{x, y}, direction: Up}
}

func isValidPosition(p Point, grid [][]string) bool {
	return p.x >= 0 && p.y >= 0 && p.x < len(grid) && p.y < len(grid[0])
}

func part1Helper(start Soldier, grid [][]string, res *int) {
	// Using bit flags to track visited directions at each point
	// Each direction is represented by a single bit in uint8:
	//   Up    (0): 00000001  (1 << 0)
	//   Right (1): 00000010  (1 << 1)
	//   Down  (2): 00000100  (1 << 2)
	//   Left  (3): 00001000  (1 << 3)
	// Benefits:
	// - Uses single byte instead of array of strings
	// - Fast bitwise operations for checking/marking visited
	// - Much lower memory usage
	visited := make(map[Point]uint8)

	for {
		next := Point{
			x: start.location.x + directionVectors[start.direction].x,
			y: start.location.y + directionVectors[start.direction].y,
		}

		if !isValidPosition(next, grid) {
			break
		}

		// Create flag for current direction by shifting 1 left by direction value
		// e.g., for Right (1): 1 << 1 = 00000010
		dirFlag := uint8(1 << start.direction)

		switch grid[next.x][next.y] {
		case ".":
			grid[next.x][next.y] = "X"
			*res++
			// Mark direction as visited using OR operation (|=)
			// e.g., 00000001 |= 00000010 = 00000011 (marks both dirs as visited)
			visited[start.location] |= dirFlag
			start.location = next
		case "X":
			// Check if direction was visited using AND operation (&)
			// e.g., 00000011 & 00000010 != 0 means Right was visited
			if visited[start.location]&dirFlag != 0 {
				panic("Infinite loop")
			}
			visited[start.location] |= dirFlag
			start.location = next
		case "#":
			start.direction = turnMap[start.direction]
		}
	}
}

func (s Solutions) Part1(file *os.File) any {
	input := make(chan string)
	go utils.ReadFileByLineIntoChannel(file, input)
	grid, start := convert(input)
	res := 1
	grid[start.location.x][start.location.y] = "X"
	part1Helper(start, grid, &res)
	return res
}

func tryPath(start Soldier, grid [][]string, res *int) (success bool) {
	defer func() {
		if r := recover(); r != nil {
			success = false
		}
	}()
	part1Helper(start, grid, res)
	return true
}

func (s Solutions) Part2(file *os.File) any {
	input := make(chan string)
	go utils.ReadFileByLineIntoChannel(file, input)
	grid, start := convert(input)

	rows, cols := len(grid), len(grid[0])
	validPositions := make([]Point, 0, rows*cols)

	grid[start.location.x][start.location.y] = "X"
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == "." {
				validPositions = append(validPositions, Point{i, j})
			}
		}
	}

	workerCount := runtime.NumCPU()
	chunkSize := (len(validPositions) + workerCount - 1) / workerCount

	var res atomic.Int64
	var wg sync.WaitGroup

	for i := 0; i < len(validPositions); i += chunkSize {
		wg.Add(1)
		end := i + chunkSize
		if end > len(validPositions) {
			end = len(validPositions)
		}

		go func(positions []Point) {
			defer wg.Done()
			gridCopy := make([][]string, rows)
			for k := range gridCopy {
				gridCopy[k] = make([]string, cols)
			}

			for _, pos := range positions {
				for k := range grid {
					copy(gridCopy[k], grid[k])
				}
				gridCopy[pos.x][pos.y] = "#"

				dummyRes := 1
				startCopy := Soldier{
					location:  Point{start.location.x, start.location.y},
					direction: Up,
				}

				if !tryPath(startCopy, gridCopy, &dummyRes) {
					res.Add(1)
				}
			}
		}(validPositions[i:end])
	}

	wg.Wait()
	return res.Load()
}

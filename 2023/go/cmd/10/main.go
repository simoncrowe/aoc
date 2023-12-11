package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"

	"github.com/simoncrowe/aoc/2023/go/internal/mathutil"
)

type Ground struct{}

type Pipe struct {
	inA, inB mathutil.Vec2
	sym      string
}

type Traversable interface {
	isOpen(mathutil.Vec2) bool
	runsAlongY() bool
	enclosable() bool
}

func (g Ground) isOpen(inDir mathutil.Vec2) bool { return false }
func (g Ground) runsAlongY() bool                { return false }
func (g Ground) enclosable() bool                { return true }

func (p Pipe) isOpen(inDir mathutil.Vec2) bool {
	return inDir == p.inA || inDir == p.inB
}

func (p Pipe) runsAlongY() bool {
	return p.inA.Y != 0 || p.inB.Y != 0
}

func (p Pipe) getOutDir(inDir mathutil.Vec2) mathutil.Vec2 {
	var out mathutil.Vec2
	if inDir == p.inA {
		out = p.inB.Flip()
	} else if inDir == p.inB {
		out = p.inA.Flip()
	} else {
		log.Fatalln("Pipe ", p, " not open to dir ", inDir)
	}
	return out
}

func (p Pipe) enclosable() bool { return false }

func main() {
	lines := []string{}
	file, err := os.Open("../input/10-input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	slices.Reverse(lines)

	grid, start := makeTraversable(lines)
	loopLocs, loopPipes := walkLoop(grid, start)
	fmt.Println("Part 1 answer :", (len(loopPipes)+1)/2)

	enclosedLocs, enclosedObjs := getEnclosed(loopLocs, loopPipes, grid)
	log.Print(enclosedLocs)
	log.Print(enclosedObjs)
	fmt.Println("Part 2 answer :", len(enclosedLocs))
}

func walkLoop(grid [][]Traversable, start mathutil.Vec2) ([]mathutil.Vec2, []Pipe) {
	pipes := []Pipe{}
	locs := []mathutil.Vec2{}
	dir := findDir(grid, start)
	loc := start.Add(dir)
	for loc != start {
		locs = append(locs, loc)
		pipe := grid[loc.Y][loc.X].(Pipe)
		log.Print("pipe at loc ", loc, ": ", pipe)
		pipes = append(pipes, pipe)
		dir = pipe.getOutDir(dir)
		loc = loc.Add(dir)
		log.Print("loc: ", loc, "; dir: ", dir)
	}
	return locs, pipes
}

func getEnclosed(loopLocs []mathutil.Vec2, loopPipes []Pipe, grid [][]Traversable) ([]mathutil.Vec2, []Traversable) {
	enclosedLocs := []mathutil.Vec2{}
	enclosedObjs := []Traversable{}
	height := len(grid)
	width := len(grid[0])
	insideLoop := false
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			obj := grid[y][x]
			loc := mathutil.Vec2{x, y}
			if obj.enclosable() && insideLoop {
				enclosedLocs = append(enclosedLocs, loc)
			}
			if slices.Contains(loopLocs, loc) && obj.runsAlongY() {
				insideLoop = !insideLoop
			}
		}
	}
	return enclosedLocs, enclosedObjs
}

func findDir(grid [][]Traversable, loc mathutil.Vec2) mathutil.Vec2 {
	dirs := []mathutil.Vec2{mathutil.Vec2{1, 0}, mathutil.Vec2{0, 1}, mathutil.Vec2{-1, 0}, mathutil.Vec2{0, -1}}
	var found mathutil.Vec2
	for _, dir := range dirs {
		x, y := loc.X+dir.X, loc.Y+dir.Y
		if x < 0 || x >= len(grid[0]) || y < 0 || y >= len(grid) {
			continue
		}
		if grid[y][x].isOpen(dir) {
			found = dir
			break
		}
	}
	return found
}

func makeTraversable(lines []string) ([][]Traversable, mathutil.Vec2) {
	width := len(lines[0])
	height := len(lines)
	data := make([][]Traversable, height)
	for i := range data {
		data[i] = make([]Traversable, width)
	}
	var startLoc mathutil.Vec2

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			switch sym := string(lines[y][x]); sym {
			case "|":
				data[y][x] = Pipe{mathutil.Vec2{0, 1}, mathutil.Vec2{0, -1}, sym}
			case "-":
				data[y][x] = Pipe{mathutil.Vec2{1, 0}, mathutil.Vec2{-1, 0}, sym}
			case "L":
				data[y][x] = Pipe{mathutil.Vec2{0, -1}, mathutil.Vec2{-1, 0}, sym}
			case "J":
				data[y][x] = Pipe{mathutil.Vec2{0, -1}, mathutil.Vec2{1, 0}, sym}
			case "7":
				data[y][x] = Pipe{mathutil.Vec2{0, 1}, mathutil.Vec2{1, 0}, sym}
			case "F":
				data[y][x] = Pipe{mathutil.Vec2{0, 1}, mathutil.Vec2{-1, 0}, sym}
			case ".":
				data[y][x] = Ground{}
			case "S":
				startLoc = mathutil.Vec2{x, y}
				log.Print("Start loc: ", startLoc)
			}
		}
	}
	return data, startLoc
}

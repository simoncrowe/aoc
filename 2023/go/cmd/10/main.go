package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"

	"github.com/EliCDavis/vector/vector2"
	"github.com/simoncrowe/aoc/2023/go/internal/mathutil"
)

type Ground struct{ sym string }

type Pipe struct {
	inA, inB vector2.Vector[int]
	sym      string
}

type Traversable interface {
	isOpen(vector2.Vector[int]) bool
	getSym() string
}

func (g Ground) isOpen(inDir vector2.Vector[int]) bool { return false }
func (g Ground) getSym() string                        { return g.sym }

func (p Pipe) isOpen(inDir vector2.Vector[int]) bool {
	return inDir == p.inA || inDir == p.inB
}

func (p Pipe) getOutDir(inDir vector2.Vector[int]) vector2.Vector[int] {
	var out vector2.Vector[int]
	if inDir == p.inA {
		out = p.inB.Flip()
	} else if inDir == p.inB {
		out = p.inA.Flip()
	} else {
		log.Fatalln("Pipe ", p, " not open to dir ", inDir)
	}
	return out
}

func (p Pipe) getSym() string { return p.sym }

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

	loopSegs := buildLineSegs(loopLocs)
	enclosed := countEnclosed(loopSegs, loopLocs, loopPipes, grid)
	fmt.Println("Part 2 answer :", enclosed)
}

func walkLoop(grid [][]Traversable, start vector2.Vector[int]) ([]vector2.Vector[int], []Pipe) {
	startPipe := grid[start.Y()][start.X()].(Pipe)
	dir := startPipe.inA.Flip()
	pipes := []Pipe{startPipe}
	locs := []vector2.Vector[int]{start}
	loc := start.Add(dir)
	for loc != start {
		locs = append(locs, loc)
		pipe := grid[loc.Y()][loc.X()].(Pipe)
		pipes = append(pipes, pipe)
		dir = pipe.getOutDir(dir)
		loc = loc.Add(dir)
	}
	return locs, pipes
}

func buildLineSegs(centres []vector2.Vector[int]) []mathutil.LineSeg {
	
}

func countEnclosed(loop []mathutil.LineSeg, loopLocs []vector2.Vector[int], loopPipes []Pipe, grid [][]Traversable) int {
	height := len(grid)
	width := len(grid[0])
	enclosed := 0
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			loc := vector2.New[int](x, y)
			obj := grid[y][x]
			if slices.Contains(loopLocs, loc) {
				log.Printf("Skipping %v (%s) as part of loop", loc, obj.getSym())
				continue
			}
		}
	}
	return enclosed 
}

func findInDirs(grid [][]Traversable, loc vector2.Vector[int]) (vector2.Vector[int], vector2.Vector[int]) {
	dirs := []vector2.Vector[int]{vector2.New[int](1, 0), vector2.New[int](0, 1), vector2.New[int](-1, 0), vector2.New[int](0, -1)}
	found := []vector2.Vector[int]{}
	for _, dir := range dirs {
		x, y := loc.X()+dir.X(), loc.Y()+dir.Y()
		if x < 0 || x >= len(grid[0]) || y < 0 || y >= len(grid) {
			continue
		}
		if grid[y][x].isOpen(dir) {
			found = append(found, dir.Flip())
		}
	}
	return found[0], found[1]
}

func makeTraversable(lines []string) ([][]Traversable, vector2.Vector[int]) {
	width := len(lines[0])
	height := len(lines)
	data := make([][]Traversable, height)
	for i := range data {
		data[i] = make([]Traversable, width)
	}
	var startLoc vector2.Vector[int]

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			switch sym := string(lines[y][x]); sym {
			case "|":
				data[y][x] = Pipe{vector2.New[int](0, 1), vector2.New[int](0, -1), sym}
			case "-":
				data[y][x] = Pipe{vector2.New[int](1, 0), vector2.New[int](-1, 0), sym}
			case "L":
				data[y][x] = Pipe{vector2.New[int](0, -1), vector2.New[int](-1, 0), sym}
			case "J":
				data[y][x] = Pipe{vector2.New[int](0, -1), vector2.New[int](1, 0), sym}
			case "7":
				data[y][x] = Pipe{vector2.New[int](0, 1), vector2.New[int](1, 0), sym}
			case "F":
				data[y][x] = Pipe{vector2.New[int](0, 1), vector2.New[int](-1, 0), sym}
			case ".":
				data[y][x] = Ground{sym}
			case "S":
				startLoc = vector2.New[int](x, y)
			}
		}
	}
	inA, inB := findInDirs(data, startLoc)
	data[startLoc.Y()][startLoc.X()] = Pipe{inA, inB, "S"}
	return data, startLoc
}

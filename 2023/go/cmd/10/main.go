package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
)

type Vec2 struct {
	x, y int
}

func (a Vec2) add(b Vec2) Vec2 {
	return Vec2{a.x + b.x, a.y + b.y}
}

func (v Vec2) flip() Vec2 {
	return Vec2{v.x - (2 * v.x), v.y - (2 * v.y)}
}

type Ground struct{}

type Pipe struct {
	inA, inB Vec2
	sym      string
}

type Traversable interface {
	isOpen(Vec2) bool
}

func (p Pipe) isOpen(inDir Vec2) bool {
	return inDir == p.inA || inDir == p.inB
}

func (g Ground) isOpen(inDir Vec2) bool {
	return false
}

func (p Pipe) exit(inDir Vec2) Vec2 {
	var out Vec2
	if inDir == p.inA {
		out = p.inB.flip()
	} else if inDir == p.inB {
		out = p.inA.flip()
	} else {
		log.Fatalln("Pipe ", p, " not open to dir ", inDir)
	}
	return out
}

func (g Ground) CanEnter(inDir Vec2) bool {
	return false
}

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
	cycle := walk(grid, start)

	fmt.Println("Part 1 answer :", (len(cycle)+1)/2)
}

func walk(grid [][]Traversable, start Vec2) []Pipe {
	cycle := []Pipe{}
	dir := findDir(grid, start)
	loc := start.add(dir)
	for loc != start {
		pipe := grid[loc.y][loc.x].(Pipe)
		log.Print("pipe at loc ", loc, ": ", pipe)
		cycle = append(cycle, pipe)
		dir = pipe.exit(dir)
		loc = loc.add(dir)
		log.Print("loc: ", loc, "; dir: ", dir)
	}
	return cycle
}

func findDir(grid [][]Traversable, loc Vec2) Vec2 {
	dirs := []Vec2{Vec2{1, 0}, Vec2{0, 1}, Vec2{-1, 0}, Vec2{0, -1}}
	var found Vec2
	for _, dir := range dirs {
		x, y := loc.x+dir.x, loc.y+dir.y
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

func makeTraversable(lines []string) ([][]Traversable, Vec2) {
	width := len(lines[0])
	height := len(lines)
	data := make([][]Traversable, height)
	for i := range data {
		data[i] = make([]Traversable, width)
	}
	var startLoc Vec2

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			switch sym := string(lines[y][x]); sym {
			case "|":
				data[y][x] = Pipe{Vec2{0, 1}, Vec2{0, -1}, sym}
			case "-":
				data[y][x] = Pipe{Vec2{1, 0}, Vec2{-1, 0}, sym}
			case "L":
				data[y][x] = Pipe{Vec2{0, -1}, Vec2{-1, 0}, sym}
			case "J":
				data[y][x] = Pipe{Vec2{0, -1}, Vec2{1, 0}, sym}
			case "7":
				data[y][x] = Pipe{Vec2{0, 1}, Vec2{1, 0}, sym}
			case "F":
				data[y][x] = Pipe{Vec2{0, 1}, Vec2{-1, 0}, sym}
			case ".":
				data[y][x] = Ground{}
			case "S":
				startLoc = Vec2{x, y}
				log.Print("Start loc: ", startLoc)
			}
		}
	}
	return data, startLoc
}
